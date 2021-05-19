package main

import (
	"flag"
	"os"
	"reflect"

	"go.uber.org/zap"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	argoeventbus "github.com/argoproj-labs/argo-eventbus"
	"github.com/argoproj-labs/argo-eventbus/controllers/eventbus"
	eventbusv1alpha1 "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1"
)

const (
	natsStreamingEnvVar       = "NATS_STREAMING_IMAGE"
	natsMetricsExporterEnvVar = "NATS_METRICS_EXPORTER_IMAGE"
)

var (
	namespaced       bool
	managedNamespace string
)

func init() {
	flag.BoolVar(&namespaced, "namespaced", false, "run the controller as namespaced mode")
	flag.StringVar(&managedNamespace, "managed-namespace", os.Getenv("NAMESPACE"), "namespace that controller watches, default to the installation namespace")
	flag.Parse()
}

func main() {
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = []string{"stdout"}
	l, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := l.Named("argo-eventbus").Sugar()
	natsStreamingImage, defined := os.LookupEnv(natsStreamingEnvVar)
	if !defined {
		logger.Fatalf("required environment variable '%s' not defined", natsStreamingEnvVar)
	}
	natsMetricsImage, defined := os.LookupEnv(natsMetricsExporterEnvVar)
	if !defined {
		logger.Fatalf("required environment variable '%s' not defined", natsMetricsExporterEnvVar)
	}
	opts := ctrl.Options{
		MetricsBindAddress:     ":7777",
		HealthProbeBindAddress: ":8081",
	}
	if namespaced {
		opts.Namespace = managedNamespace
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), opts)
	if err != nil {
		logger.Fatalw("unable to get a controller-runtime manager", zap.Error(err))
	}

	// Readyness probe
	if err := mgr.AddReadyzCheck("readiness", healthz.Ping); err != nil {
		logger.Fatalw("unable add a readiness check", zap.Error(err))
	}

	// Liveness probe
	if err := mgr.AddHealthzCheck("liveness", healthz.Ping); err != nil {
		logger.Fatalw("unable add a health check", zap.Error(err))
	}

	if err := eventbusv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		logger.Fatalw("unable to add scheme", zap.Error(err))
	}

	// A controller with DefaultControllerRateLimiter
	c, err := controller.New(eventbus.ControllerName, mgr, controller.Options{
		Reconciler: eventbus.NewReconciler(mgr.GetClient(), mgr.GetScheme(), natsStreamingImage, natsMetricsImage, logger),
	})
	if err != nil {
		logger.Fatalw("unable to set up individual controller", zap.Error(err))
	}

	// Watch EventBus and enqueue EventBus object key
	if err := c.Watch(&source.Kind{Type: &eventbusv1alpha1.EventBus{}}, &handler.EnqueueRequestForObject{},
		predicate.Or(
			predicate.GenerationChangedPredicate{},
			// TODO: change to use LabelChangedPredicate with controller-runtime v0.8
			predicate.Funcs{
				UpdateFunc: func(e event.UpdateEvent) bool {
					if e.ObjectOld == nil {
						return false
					}
					if e.ObjectNew == nil {
						return false
					}
					return !reflect.DeepEqual(e.ObjectNew.GetLabels(), e.ObjectOld.GetLabels())
				}},
		)); err != nil {
		logger.Fatalw("unable to watch EventBus", zap.Error(err))
	}

	// Watch ConfigMaps and enqueue owning EventBus key
	if err := c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{OwnerType: &eventbusv1alpha1.EventBus{}, IsController: true}, predicate.GenerationChangedPredicate{}); err != nil {
		logger.Fatalw("unable to watch ConfigMaps", zap.Error(err))
	}

	// Watch Secrets and enqueue owning EventBus key
	if err := c.Watch(&source.Kind{Type: &corev1.Secret{}}, &handler.EnqueueRequestForOwner{OwnerType: &eventbusv1alpha1.EventBus{}, IsController: true}, predicate.GenerationChangedPredicate{}); err != nil {
		logger.Fatalw("unable to watch Secrets", zap.Error(err))
	}

	// Watch StatefulSets and enqueue owning EventBus key
	if err := c.Watch(&source.Kind{Type: &appv1.StatefulSet{}}, &handler.EnqueueRequestForOwner{OwnerType: &eventbusv1alpha1.EventBus{}, IsController: true}, predicate.GenerationChangedPredicate{}); err != nil {
		logger.Fatalw("unable to watch StatefulSets", zap.Error(err))
	}

	// Watch Services and enqueue owning EventBus key
	if err := c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{OwnerType: &eventbusv1alpha1.EventBus{}, IsController: true}, predicate.GenerationChangedPredicate{}); err != nil {
		logger.Fatalw("unable to watch Services", zap.Error(err))
	}

	logger.Infow("starting eventbus controller", "version", argoeventbus.GetVersion())
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		logger.Fatalw("unable to run eventbus controller", zap.Error(err))
	}
}
