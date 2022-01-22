/*
Copyright 2022 Alessio G. Baroni.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DekerisSpec defines the desired state of Dekeris
type DekerisSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Dekeris. Edit dekeris_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DekerisStatus defines the observed state of Dekeris
type DekerisStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Dekeris is the Schema for the dekeris API
type Dekeris struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DekerisSpec   `json:"spec,omitempty"`
	Status DekerisStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DekerisList contains a list of Dekeris
type DekerisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dekeris `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dekeris{}, &DekerisList{})
}
