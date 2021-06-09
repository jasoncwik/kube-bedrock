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

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BedrockServerSpec defines the desired state of BedrockServer
// +kubebuilder:printcolumn:name="Name",type=string,JSONPath=`.spec.name`
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`
// +kubebuilder:printcolumn:name="Run",type=string,JSONPath=`.spec.run`
// +kubebuilder:printcolumn:name="StorageClass",type=string,JSONPath=`.spec.worldsStorageClass`
// +kubebuilder:printcolumn:name="Size",type=string,JSONPath=`.spec.worldsSize `
type BedrockServerSpec struct {
	// Name of the server
	Name               string            `json:"name,omitempty"`

	// The container image (version) to execute
	Image              v1.ContainerImage `json:"image"`

	// Whether the server should be running. Can be used to take servers offline.
	// Setting this to False will scale the StatefulSet to zero.
	Run                bool              `json:"run"`

	// The name of the StorageClass to use for the PV to hold the
	// "Worlds" directory. This is where your server state is stored. If
	// not specified the default storage class will be used.
	WorldsStorageClass string            `json:"worldsStorageClass,omitempty"`

	// Size of the PV to create. Defaults to 1Gi
	WorldsSize         resource.Quantity `json:"worldsSize,omitempty"`

	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	// Type of Service to create to expose the server. Defaults to
	// ClusterIP, something that is not particularly useful. Recommend
	// using NodePort or LoadBalancer.
	ServiceType v1.ServiceType `json:"serviceType,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// Port to expose the service on. Default is 19132 for LoadBalancer
	// or ClusterIP and autoselected for NodePort.
	ServicePort int32	`json:"servicePort,omitempty"`
}

// BedrockServerStatus defines the observed state of BedrockServer
type BedrockServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BedrockServer is the Schema for the bedrockservers API
type BedrockServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BedrockServerSpec   `json:"spec,omitempty"`
	Status BedrockServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BedrockServerList contains a list of BedrockServer
type BedrockServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BedrockServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BedrockServer{}, &BedrockServerList{})
}
