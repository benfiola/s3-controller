package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true
type BucketSpec struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:generate=true
type BucketStatus struct {
	Conditions         []metav1.Condition `json:"conditions,omitempty"`
	LastReconciledTime *metav1.Time       `json:"lastReconciledTime,omitempty"`
	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Bucket `json:"items"`
}
