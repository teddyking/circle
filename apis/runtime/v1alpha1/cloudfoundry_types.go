/*

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

// CloudFoundrySpec defines the desired state of CloudFoundry
type CloudFoundrySpec struct {
	API               string                `json:"api,omitempty"`
	AuthInfo          *CloudFoundryAuthInfo `json:"authInfo,omitempty"`
	SkipSSLValidation bool                  `json:"skipSSLValidation,omitempty"`
}

// CloudFoundryStatus defines the observed state of CloudFoundry
type CloudFoundryStatus struct {
	LastReconcileTime *metav1.Time `json:"lastReconcileTime,omitempty"`
}

// CloudFoundryAuthInfo provides config for authenticating to the CloudFoundry
type CloudFoundryAuthInfo struct {
	SecretRef *ObjectReference `json:"secretRef,omitempty"`
}

// ObjectReference contains enough information to let you locate the
// referenced object.
type ObjectReference struct {
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
	// Name of the referent.
	Name string `json:"name,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=cf

// CloudFoundry is the Schema for the cloudfoundries API
type CloudFoundry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudFoundrySpec   `json:"spec,omitempty"`
	Status CloudFoundryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudFoundryList contains a list of CloudFoundry
type CloudFoundryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudFoundry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudFoundry{}, &CloudFoundryList{})
}
