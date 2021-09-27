/*
Copyright 2021 The Crossplane Authors.

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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// KeyValueParameters are the configurable fields of a KeyValue.
type KeyValueParameters struct {
	// Endpoint of the App Configuration.
	Endpoint string `json:"endpoint"`

	// Key of the KeyValue.
	// +immutable
	Key string `json:"key"`

	// Label of the KeyValue.
	// +immutable
	Label string `json:"label"`

	// Value of the KeyValue
	Value string `json:"value"`

	// Content-Type of the KeyValue.
	// +optional
	ContentType *string `json:"contentType,omitempty"`

	// Whether this KeyValue is locked.
	// +optional
	Locked *bool `json:"locked,omitempty"`

	// Tags of the KeyValue resource.
	// +optional
	Tags *map[string]string `json:"tags,omitempty"`
}

// A KeyValueSpec defines the desired state of a KeyValue.
type KeyValueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KeyValueParameters `json:"forProvider"`
}

// KeyValueObservation represents the observed state of the Secret object in Azure.
type KeyValueObservation struct {
	// Time that the KeyValue was last modified.
	LastModified string `json:"last_modified,omitempty"`

	// ETag of the KeyValue.
	Etag string `json:"etag,omitempty"`
}

// A KeyValueStatus represents the observed state of a KeyValue.
type KeyValueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KeyValueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A KeyValue is a managed resource that represents a App Configuration KeyValue pai.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KeyValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyValueSpec   `json:"spec"`
	Status KeyValueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyValueList contains a list of KeyValue
type KeyValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyValue `json:"items"`
}
