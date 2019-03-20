package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MultipleNamespaceFederationSpec defines the desired state of MultipleNamespaceFederation
// +k8s:openapi-gen=true
type MultipleNamespaceFederationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Clusters          []string              `json:"clusters,omitempty"`
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector"`
}

// MultipleNamespaceFederationStatus defines the observed state of MultipleNamespaceFederation
// +k8s:openapi-gen=true
type MultipleNamespaceFederationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MultipleNamespaceFederation is the Schema for the multiplenamespacefederations API
// +k8s:openapi-gen=true
type MultipleNamespaceFederation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultipleNamespaceFederationSpec   `json:"spec,omitempty"`
	Status MultipleNamespaceFederationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MultipleNamespaceFederationList contains a list of MultipleNamespaceFederation
type MultipleNamespaceFederationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MultipleNamespaceFederation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MultipleNamespaceFederation{}, &MultipleNamespaceFederationList{})
}