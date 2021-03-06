package installer

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1"
)

// Installer is an interface for event bus installation
type Installer interface {
	Install(ctx context.Context) (*v1alpha1.BusConfig, error)
	// Uninsall only needs to handle those resources not cascade deleted.
	// For example, undeleted PVCs not automatically deleted when deleting a StatefulSet
	Uninstall(ctx context.Context) error
}

// Install function installs the event bus
func Install(ctx context.Context, eventBus *v1alpha1.EventBus, client client.Client, natsStreamingImage, natsMetricsImage string, logger *zap.SugaredLogger) error {
	installer, err := getInstaller(eventBus, client, natsStreamingImage, natsMetricsImage, logger)
	if err != nil {
		logger.Errorw("failed to an installer", zap.Error(err))
		return err
	}
	busConfig, err := installer.Install(ctx)
	if err != nil {
		logger.Errorw("installation error", zap.Error(err))
		return err
	}
	eventBus.Status.Config = *busConfig
	return nil
}

// GetInstaller returns Installer implementation
func getInstaller(eventBus *v1alpha1.EventBus, client client.Client, natsStreamingImage, natsMetricsImage string, logger *zap.SugaredLogger) (Installer, error) {
	if nats := eventBus.Spec.NATS; nats != nil {
		if nats.Exotic != nil {
			return NewExoticNATSInstaller(eventBus, logger), nil
		} else if nats.Native != nil {
			return NewNATSInstaller(client, eventBus, natsStreamingImage, natsMetricsImage, getLabels(eventBus), logger), nil
		}
	}
	return nil, errors.New("invalid eventbus spec")
}

func getLabels(bus *v1alpha1.EventBus) map[string]string {
	return map[string]string{
		"controller":    "eventbus-controller",
		"eventbus-name": bus.Name,
		"owner-name":    bus.Name,
	}
}

// Uninstall function will be run before the EventBus object is deleted,
// usually it could be used to uninstall the extra resources who would not be cleaned
// up when an EventBus is deleted. Most of the time this is not needed as all
// the dependency resources should have been deleted by owner references cascade
// deletion, but things like PVC created by StatefulSet need to be cleaned up
// separately.
//
// It could also be used to check if the EventBus object can be safely deleted.
func Uninstall(ctx context.Context, eventBus *v1alpha1.EventBus, client client.Client, natsStreamingImage, natsMetricsImage string, logger *zap.SugaredLogger) error {
	installer, err := getInstaller(eventBus, client, natsStreamingImage, natsMetricsImage, logger)
	if err != nil {
		logger.Errorw("failed to get an installer", zap.Error(err))
		return err
	}
	return installer.Uninstall(ctx)
}
