/*
Copyright 2023.

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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BoxSpec defines the desired state of Box
type BoxSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Length is the length of the box
	Length int `json:"length,omitempty"`
	// Width is the width of the box
	Width int `json:"width,omitempty"`
	// Height is the height of the box
	Height int `json:"height,omitempty"`
	// Weight is the weight of the box
	Weight int `json:"weight,omitempty"`
	// +optional
	// +kubebuilder:default:="empty"
	// Content is a short description of the content of the box
	Content string `json:"content,omitempty"`
}

// BoxStatus defines the observed state of Box
type BoxStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Box is the Schema for the boxes API
type Box struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BoxSpec   `json:"spec,omitempty"`
	Status BoxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BoxList contains a list of Box
type BoxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Box `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Box{}, &BoxList{})
}
