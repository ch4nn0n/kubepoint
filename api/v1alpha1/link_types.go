/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LinkSpec defines the desired state of Link
type LinkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Link name
	Name string `json:"name"`
	// Subtitle
	// +optional
	Subtitle string `json:"subtitle"`
	// Target url
	Url string `json:"url"`
	// Logo
	// +optional
	Logo string `json:"logo"`
	// Icon
	// +optional
	Icon string `json:"icon"`
	// Tag
	// +optional
	Tag string `json:"tag"`
	// Tag style
	// Can be set to any of the bulma modifiers, thought you probably want to use:
	// - is-info (blue)
	// - is-success (green)
	// - is-warning (yellow)
	// - is-danger (red)
	// https://bulma.io/documentation/overview/modifiers/
	// +optional
	TagStyle string `json:"tagStyle"`
	// Type
	// Loads a specific component that provides extra features.
	// MUST MATCH a file name (without file extension) available in `https://github.com/bastienwirtz/homer/tree/main/src/components/services`
	// +optional
	Type string `json:"type"`
	// Target HTML target tag
	// +optional
	Target string `json:"target"`
	// Background
	// to set color for the card directly
	// + optional
	Background string `json:"background"`
}

// LinkStatus defines the observed state of Link
type LinkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Link is the Schema for the links API
type Link struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinkSpec   `json:"spec,omitempty"`
	Status LinkStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LinkList contains a list of Link
type LinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Link `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Link{}, &LinkList{})
}
