module github.com/argoproj-labs/argo-eventbus

go 1.16

replace k8s.io/api => k8s.io/api v0.19.6

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.6

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.8-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.19.6

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.6

replace k8s.io/client-go => k8s.io/client-go v0.19.6

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.6

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.6

replace k8s.io/code-generator => k8s.io/code-generator v0.19.9-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.19.6

replace k8s.io/controller-manager => k8s.io/controller-manager v0.19.11-rc.0

replace k8s.io/cri-api => k8s.io/cri-api v0.19.9-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.6

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.6

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.6

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.6

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.6

replace k8s.io/kubectl => k8s.io/kubectl v0.19.6

replace k8s.io/kubelet => k8s.io/kubelet v0.19.6

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.6

replace k8s.io/metrics => k8s.io/metrics v0.19.6

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.6

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.6

replace k8s.io/sample-controller => k8s.io/sample-controller v0.19.6

require (
	github.com/ahmetb/gen-crd-api-reference-docs v0.2.0
	github.com/emicklei/go-restful v2.12.0+incompatible // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/logr v0.3.0 // indirect
	github.com/go-openapi/spec v0.20.2 // indirect
	github.com/go-swagger/go-swagger v0.25.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/smartystreets/assertions v0.0.0-20190401211740-f487f9de1cd3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	golang.org/x/mod v0.3.1-0.20200828183125-ce943fd02449 // indirect
	golang.org/x/net v0.0.0-20210326060303-6b1517762897 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	k8s.io/api v0.19.6
	k8s.io/apiextensions-apiserver v0.19.2 // indirect
	k8s.io/apimachinery v0.19.6
	k8s.io/code-generator v0.19.6
	k8s.io/gengo v0.0.0-20200428234225-8167cfdcfc14
	k8s.io/klog v0.3.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/controller-tools v0.4.1
)
