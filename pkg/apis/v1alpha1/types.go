package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventBus is the definition of a eventbus resource
// +genclient
// +kubebuilder:resource:singular=eventbus,shortName=eb
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type EventBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Spec              EventBusSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	// +optional
	Status EventBusStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// EventBusList is the list of eventbus resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type EventBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	Items []EventBus `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// EventBusSpec refers to specification of eventbus resource
type EventBusSpec struct {
	// NATS eventbus
	NATS *NATSBus `json:"nats,omitempty" protobuf:"bytes,1,opt,name=nats"`
}

// EventBusStatus holds the status of the eventbus resource
type EventBusStatus struct {
	Status `json:",inline" protobuf:"bytes,1,opt,name=status"`
	// Config holds the fininalized configuration of EventBus
	Config BusConfig `json:"config,omitempty" protobuf:"bytes,2,opt,name=config"`
}

// NATSBus holds the NATS eventbus information
type NATSBus struct {
	// Native means to bring up a native NATS service
	Native *NativeStrategy `json:"native,omitempty" protobuf:"bytes,1,opt,name=native"`
	// Exotic holds an exotic NATS config
	Exotic *NATSConfig `json:"exotic,omitempty" protobuf:"bytes,2,opt,name=exotic"`
}

// AuthStrategy is the auth strategy of native nats installaion
type AuthStrategy string

// possible auth strategies
var (
	AuthStrategyNone  AuthStrategy = "none"
	AuthStrategyToken AuthStrategy = "token"
)

// NativeStrategy indicates to install a native NATS service
type NativeStrategy struct {
	// Size is the NATS StatefulSet size
	Replicas int32         `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	Auth     *AuthStrategy `json:"auth,omitempty" protobuf:"bytes,2,opt,name=auth,casttype=AuthStrategy"`
	// Deprecated, use Affinity instead, will be removed in v1.5
	DeprecatedAntiAffinity bool `json:"antiAffinity,omitempty" protobuf:"varint,3,opt,name=antiAffinity"`
	// +optional
	Persistence *PersistenceStrategy `json:"persistence,omitempty" protobuf:"bytes,4,opt,name=persistence"`
	// ContainerTemplate contains customized spec for NATS container
	// +optional
	ContainerTemplate *ContainerTemplate `json:"containerTemplate,omitempty" protobuf:"bytes,5,opt,name=containerTemplate"`
	// MetricsContainerTemplate contains customized spec for metrics container
	// +optional
	MetricsContainerTemplate *ContainerTemplate `json:"metricsContainerTemplate,omitempty" protobuf:"bytes,6,opt,name=metricsContainerTemplate"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty" protobuf:"bytes,8,rep,name=tolerations"`
	// Metadata sets the pods's metadata, i.e. annotations and labels
	Metadata *Metadata `json:"metadata,omitempty" protobuf:"bytes,9,opt,name=metadata"`
	// SecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,10,opt,name=securityContext"`
	// Max Age of existing messages, i.e. "72h", “4h35m”
	// +optional
	MaxAge *string `json:"maxAge,omitempty" protobuf:"bytes,11,opt,name=maxAge"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,12,rep,name=imagePullSecrets"`
	// ServiceAccountName to apply to NATS StatefulSet
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty" protobuf:"bytes,13,opt,name=serviceAccountName"`
	// If specified, indicates the EventSource pod's priority. "system-node-critical"
	// and "system-cluster-critical" are two special keywords which indicate the
	// highest priorities with the former being the highest priority. Any other
	// name must be defined by creating a PriorityClass object with that name.
	// If not specified, the pod priority will be default or zero if there is no
	// default.
	// More info: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/
	// +optional
	PriorityClassName string `json:"priorityClassName,omitempty" protobuf:"bytes,14,opt,name=priorityClassName"`
	// The priority value. Various system components use this field to find the
	// priority of the EventSource pod. When Priority Admission Controller is enabled,
	// it prevents users from setting this field. The admission controller populates
	// this field from PriorityClassName.
	// The higher the value, the higher the priority.
	// More info: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/
	// +optional
	Priority *int32 `json:"priority,omitempty" protobuf:"bytes,15,opt,name=priority"`
	// The pod's scheduling constraints
	// More info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty" protobuf:"bytes,16,opt,name=affinity"`
}

// ContainerTemplate defines customized spec for a container
type ContainerTemplate struct {
	Resources corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,1,opt,name=resources"`
}

// GetReplicas return the replicas of statefulset
func (in *NativeStrategy) GetReplicas() int {
	return int(in.Replicas)
}

// PersistenceStrategy defines the strategy of persistence
type PersistenceStrategy struct {
	// Name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty" protobuf:"bytes,1,opt,name=storageClassName"`
	// Available access modes such as ReadWriteOnce, ReadWriteMany
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes
	// +optional
	AccessMode *corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty" protobuf:"bytes,2,opt,name=accessMode,casttype=k8s.io/api/core/v1.PersistentVolumeAccessMode"`
	// Volume size, e.g. 10Gi
	VolumeSize *apiresource.Quantity `json:"volumeSize,omitempty" protobuf:"bytes,3,opt,name=volumeSize"`
}

// BusConfig has the finalized configuration for EventBus
type BusConfig struct {
	NATS *NATSConfig `json:"nats,omitempty" protobuf:"bytes,1,opt,name=nats"`
}

// NATSConfig holds the config of NATS
type NATSConfig struct {
	// NATS streaming url
	URL string `json:"url,omitempty" protobuf:"bytes,1,opt,name=url"`
	// Cluster ID for nats streaming
	ClusterID *string `json:"clusterID,omitempty" protobuf:"bytes,2,opt,name=clusterID"`
	// Auth strategy, default to AuthStrategyNone
	// +optional
	Auth *AuthStrategy `json:"auth,omitempty" protobuf:"bytes,3,opt,name=auth,casttype=AuthStrategy"`
	// Secret for auth
	// +optional
	AccessSecret *corev1.SecretKeySelector `json:"accessSecret,omitempty" protobuf:"bytes,4,opt,name=accessSecret"`
}

// Metadata holds the annotations and labels of an event source pod
type Metadata struct {
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,1,rep,name=annotations"`
	Labels      map[string]string `json:"labels,omitempty" protobuf:"bytes,2,rep,name=labels"`
}

const (
	// EventBusConditionDeployed has the status True when the EventBus
	// has its RestfulSet/Deployment ans service created.
	EventBusConditionDeployed ConditionType = "Deployed"
	// EventBusConditionConfigured has the status True when the EventBus
	// has its configuration ready.
	EventBusConditionConfigured ConditionType = "Configured"
)

// InitConditions sets conditions to Unknown state.
func (s *EventBusStatus) InitConditions() {
	s.InitializeConditions(EventBusConditionDeployed, EventBusConditionConfigured)
}

// MarkDeployed set the bus has been deployed.
func (s *EventBusStatus) MarkDeployed(reason, message string) {
	s.MarkTrueWithReason(EventBusConditionDeployed, reason, message)
}

// MarkDeploying set the bus is deploying
func (s *EventBusStatus) MarkDeploying(reason, message string) {
	s.MarkUnknown(EventBusConditionDeployed, reason, message)
}

// MarkDeployFailed set the bus deploy failed
func (s *EventBusStatus) MarkDeployFailed(reason, message string) {
	s.MarkFalse(EventBusConditionDeployed, reason, message)
}

// MarkConfigured set the bus configuration has been done.
func (s *EventBusStatus) MarkConfigured() {
	s.MarkTrue(EventBusConditionConfigured)
}

// MarkNotConfigured set the bus status not configured.
func (s *EventBusStatus) MarkNotConfigured(reason, message string) {
	s.MarkFalse(EventBusConditionConfigured, reason, message)
}
