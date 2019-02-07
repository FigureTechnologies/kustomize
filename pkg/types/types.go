package types

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// BlueGreenDeployment ensures that a specified number of pod replicas are running at any given time.
type BlueGreenDeployment struct {
	metav1.TypeMeta `json:",inline"`

	// If the Labels of a BlueGreenDeployment are empty, they are defaulted to
	// be the same as the Pod(s) that the BlueGreenDeployment manages.
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the specification of the desired behavior of the BlueGreenDeployment.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec BlueGreenDeploymentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status is the most recently observed status of the BlueGreenDeployment.
	// This data may be out of date by some window of time.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status BlueGreenDeploymentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func ( BlueGreenDeployment) DeepCopyObject() runtime.Object {
	panic("implement me")
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlueGreenDeploymentList is a collection of BlueGreenDeployments.
type BlueGreenDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of BlueGreenDeployments.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []BlueGreenDeployment `json:"items" protobuf:"bytes,2,rep,name=items"`
}

func ( BlueGreenDeploymentList) DeepCopyObject() runtime.Object {
	panic("implement me")
}

// BlueGreenDeploymentSpec is the specification of a BlueGreenDeployment.
type BlueGreenDeploymentSpec struct {
	// Replicas is the number of desired replicas.
	// This is a pointer to distinguish between explicit zero and unspecified.
	// Defaults to 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`

	// Minimum number of seconds for which a newly created pod should be ready
	// without any of its container crashing, for it to be considered available.
	// Defaults to 0 (pod will be considered available as soon as it is ready)
	// +optional
	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,4,opt,name=minReadySeconds"`

	// Selector is a label query over pods that should match the replica count.
	// Label keys and values that must match in order to be controlled by this replica set.
	// It must match the pod template's labels.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector *metav1.LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector"`

	// Template is the object that describes the pod that will be created if
	// insufficient replicas are detected.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// +optional
	Template v1.PodTemplateSpec `json:"template,omitempty" protobuf:"bytes,3,opt,name=template"`

	// Service which points to active bgd BlueGreenDeployment
	Service v1.Service `json:"service" protobuf:"bytes,4,opt,name=service"`
}

// BlueGreenDeploymentStatus represents the current status of a BlueGreenDeployment.
type BlueGreenDeploymentStatus struct {
	// Replicas is the most recently oberved number of replicas.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas int32 `json:"replicas" protobuf:"varint,1,opt,name=replicas"`

	// The number of pods that have labels matching the labels of the pod template of the BlueGreenDeployment.
	// +optional
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty" protobuf:"varint,2,opt,name=fullyLabeledReplicas"`

	// The number of ready replicas for this replica set.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,4,opt,name=readyReplicas"`

	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty" protobuf:"varint,5,opt,name=availableReplicas"`

	// ObservedGeneration reflects the generation of the most recently observed BlueGreenDeployment.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,3,opt,name=observedGeneration"`

	// Represents the latest available observations of a replica set's current state.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []BlueGreenDeploymentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

type BlueGreenDeploymentConditionType string

// These are valid conditions of a replica set.
const (
	// BlueGreenDeploymentReplicaFailure is added in a replica set when one of its pods fails to be created
	// due to insufficient quota, limit ranges, pod security policy, node selectors, etc. or deleted
	// due to kubelet being down or finalizers are failing.
	BlueGreenDeploymentReplicaFailure BlueGreenDeploymentConditionType = "ReplicaFailure"
)

// BlueGreenDeploymentCondition describes the state of a replica set at a certain point.
type BlueGreenDeploymentCondition struct {
	// Type of replica set condition.
	Type BlueGreenDeploymentConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=BlueGreenDeploymentConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}
