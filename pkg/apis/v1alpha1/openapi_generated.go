// +build !ignore_autogenerated

/*
Copyright 2021 Intuit, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.BusConfig":           schema_argo_eventbus_pkg_apis_v1alpha1_BusConfig(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Condition":           schema_argo_eventbus_pkg_apis_v1alpha1_Condition(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.ContainerTemplate":   schema_argo_eventbus_pkg_apis_v1alpha1_ContainerTemplate(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBus":            schema_argo_eventbus_pkg_apis_v1alpha1_EventBus(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusList":        schema_argo_eventbus_pkg_apis_v1alpha1_EventBusList(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusSpec":        schema_argo_eventbus_pkg_apis_v1alpha1_EventBusSpec(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusStatus":      schema_argo_eventbus_pkg_apis_v1alpha1_EventBusStatus(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Metadata":            schema_argo_eventbus_pkg_apis_v1alpha1_Metadata(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSBus":             schema_argo_eventbus_pkg_apis_v1alpha1_NATSBus(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSConfig":          schema_argo_eventbus_pkg_apis_v1alpha1_NATSConfig(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NativeStrategy":      schema_argo_eventbus_pkg_apis_v1alpha1_NativeStrategy(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.PersistenceStrategy": schema_argo_eventbus_pkg_apis_v1alpha1_PersistenceStrategy(ref),
		"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Status":              schema_argo_eventbus_pkg_apis_v1alpha1_Status(ref),
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_BusConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BusConfig has the finalized configuration for EventBus",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nats": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSConfig"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition contains details about resource state",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Condition type.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Condition status, True, False or Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique, this should be a short, machine understandable string that gives the reason for condition's last transition. For example, \"ImageNotFound\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_ContainerTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContainerTemplate defines customized spec for a container",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_EventBus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBus is the definition of a eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusSpec", "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBusStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_EventBusList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusList is the list of eventbus resources",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.EventBus", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_EventBusSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusSpec refers to specification of eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nats": {
						SchemaProps: spec.SchemaProps{
							Description: "NATS eventbus",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSBus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSBus"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_EventBusStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventBusStatus holds the status of the eventbus resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions are the latest available observations of a resource's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Description: "Config holds the fininalized configuration of EventBus",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.BusConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.BusConfig", "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Condition"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_Metadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Metadata holds the annotations and labels of an event source pod",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_NATSBus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NATSBus holds the NATS eventbus information",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"native": {
						SchemaProps: spec.SchemaProps{
							Description: "Native means to bring up a native NATS service",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NativeStrategy"),
						},
					},
					"exotic": {
						SchemaProps: spec.SchemaProps{
							Description: "Exotic holds an exotic NATS config",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NATSConfig", "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.NativeStrategy"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_NATSConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NATSConfig holds the config of NATS",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "NATS streaming url",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"clusterID": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster ID for nats streaming",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Description: "Auth strategy, default to AuthStrategyNone",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret for auth",
							Ref:         ref("k8s.io/api/core/v1.SecretKeySelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretKeySelector"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_NativeStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NativeStrategy indicates to install a native NATS service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Size is the NATS StatefulSet size",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"antiAffinity": {
						SchemaProps: spec.SchemaProps{
							Description: "Deprecated, use Affinity instead, will be removed in v1.5",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"persistence": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.PersistenceStrategy"),
						},
					},
					"containerTemplate": {
						SchemaProps: spec.SchemaProps{
							Description: "ContainerTemplate contains customized spec for NATS container",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.ContainerTemplate"),
						},
					},
					"metricsContainerTemplate": {
						SchemaProps: spec.SchemaProps{
							Description: "MetricsContainerTemplate contains customized spec for metrics container",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.ContainerTemplate"),
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, the pod's tolerations.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Metadata sets the pods's metadata, i.e. annotations and labels",
							Ref:         ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Metadata"),
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Description: "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.",
							Ref:         ref("k8s.io/api/core/v1.PodSecurityContext"),
						},
					},
					"maxAge": {
						SchemaProps: spec.SchemaProps{
							Description: "Max Age of existing messages, i.e. \"72h\", “4h35m”",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "name",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceAccountName to apply to NATS StatefulSet",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"priorityClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, indicates the EventSource pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default. More info: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"priority": {
						SchemaProps: spec.SchemaProps{
							Description: "The priority value. Various system components use this field to find the priority of the EventSource pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. More info: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "The pod's scheduling constraints More info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"maxMsgs": {
						SchemaProps: spec.SchemaProps{
							Description: "Maximum number of messages per channel, 0 means unlimited. Defaults to 1000000",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"maxBytes": {
						SchemaProps: spec.SchemaProps{
							Description: "Total size of messages per channel, 0 means unlimited. Defaults to 1GB",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.ContainerTemplate", "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Metadata", "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.PersistenceStrategy", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.PodSecurityContext", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_PersistenceStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PersistenceStrategy defines the strategy of persistence",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessMode": {
						SchemaProps: spec.SchemaProps{
							Description: "Available access modes such as ReadWriteOnce, ReadWriteMany https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"volumeSize": {
						SchemaProps: spec.SchemaProps{
							Description: "Volume size, e.g. 10Gi",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_argo_eventbus_pkg_apis_v1alpha1_Status(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Status is a common structure which can be used for Status field.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions are the latest available observations of a resource's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1.Condition"},
	}
}
