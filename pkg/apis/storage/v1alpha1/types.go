/*
Copyright 2018 The Crossplane Authors.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/crossplaneio/crossplane/pkg/apis/core/v1alpha1"
)

// MySQLInstanceSpec specifies the configuration of a MySQL instance.
type MySQLInstanceSpec struct {
	corev1alpha1.ResourceClaimSpec `json:",inline"`

	// mysql instance properties
	// +kubebuilder:validation:Enum=5.6,5.7
	EngineVersion string `json:"engineVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLInstance is the CRD type for abstract MySQL database instances
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="VERSION",type="string",JSONPath=".spec.engineVersion"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
type MySQLInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLInstanceSpec                `json:"spec,omitempty"`
	Status corev1alpha1.ResourceClaimStatus `json:"status,omitempty"`
}

// SetBindingPhase of this MySQLInstance.
func (i *MySQLInstance) SetBindingPhase(p corev1alpha1.BindingPhase) {
	i.Status.SetBindingPhase(p)
}

// GetBindingPhase of this MySQLInstance.
func (i *MySQLInstance) GetBindingPhase() corev1alpha1.BindingPhase {
	return i.Status.GetBindingPhase()
}

// SetConditions of this MySQLInstance.
func (i *MySQLInstance) SetConditions(c ...corev1alpha1.Condition) {
	i.Status.SetConditions(c...)
}

// SetClassReference of this MySQLInstance.
func (i *MySQLInstance) SetClassReference(r *corev1.ObjectReference) {
	i.Spec.ClassReference = r
}

// GetClassReference of this MySQLInstance.
func (i *MySQLInstance) GetClassReference() *corev1.ObjectReference {
	return i.Spec.ClassReference
}

// SetResourceReference of this MySQLInstance.
func (i *MySQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	i.Spec.ResourceReference = r
}

// GetResourceReference of this MySQLInstance.
func (i *MySQLInstance) GetResourceReference() *corev1.ObjectReference {
	return i.Spec.ResourceReference
}

// SetWriteConnectionSecretToReference of this MySQLInstance.
func (i *MySQLInstance) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	i.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this MySQLInstance.
func (i *MySQLInstance) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return i.Spec.WriteConnectionSecretToReference
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLInstanceList contains a list of MySQLInstance
type MySQLInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLInstance `json:"items"`
}

// PostgreSQLInstanceSpec specifies the configuration of this
// PostgreSQLInstance.
type PostgreSQLInstanceSpec struct {
	corev1alpha1.ResourceClaimSpec `json:",inline"`

	// postgresql instance properties
	// +kubebuilder:validation:Enum=9.6
	EngineVersion string `json:"engineVersion,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgreSQLInstance is the CRD type for abstract PostgreSQL database instances
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="VERSION",type="string",JSONPath=".spec.engineVersion"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
type PostgreSQLInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgreSQLInstanceSpec           `json:"spec,omitempty"`
	Status corev1alpha1.ResourceClaimStatus `json:"status,omitempty"`
}

// SetBindingPhase of this PostgreSQLInstance.
func (i *PostgreSQLInstance) SetBindingPhase(p corev1alpha1.BindingPhase) {
	i.Status.SetBindingPhase(p)
}

// GetBindingPhase of this PostgreSQLInstance.
func (i *PostgreSQLInstance) GetBindingPhase() corev1alpha1.BindingPhase {
	return i.Status.GetBindingPhase()
}

// SetConditions of this PostgreSQLInstance.
func (i *PostgreSQLInstance) SetConditions(c ...corev1alpha1.Condition) {
	i.Status.SetConditions(c...)
}

// SetClassReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) SetClassReference(r *corev1.ObjectReference) {
	i.Spec.ClassReference = r
}

// GetClassReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) GetClassReference() *corev1.ObjectReference {
	return i.Spec.ClassReference
}

// SetResourceReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	i.Spec.ResourceReference = r
}

// GetResourceReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) GetResourceReference() *corev1.ObjectReference {
	return i.Spec.ResourceReference
}

// SetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	i.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (i *PostgreSQLInstance) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return i.Spec.WriteConnectionSecretToReference
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgreSQLInstanceList contains a list of PostgreSQLInstance
type PostgreSQLInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgreSQLInstance `json:"items"`
}

// LocalPermissionType - Base type for LocalPermissions
type LocalPermissionType string

const (
	// ReadOnlyPermission will grant read objects in a bucket
	ReadOnlyPermission LocalPermissionType = "Read"
	// WriteOnlyPermission will grant write/delete objects in a bucket
	WriteOnlyPermission LocalPermissionType = "Write"
	// ReadWritePermission LocalPermissionType Grant both read and write permissions
	ReadWritePermission LocalPermissionType = "ReadWrite"
)

// PredefinedACL represents predefied bucket ACLs.
type PredefinedACL string

// Predefined ACLs.
const (
	ACLPrivate           PredefinedACL = "Private"
	ACLPublicRead        PredefinedACL = "PublicRead"
	ACLPublicReadWrite   PredefinedACL = "PublicReadWrite"
	ACLAuthenticatedRead PredefinedACL = "AuthenticatedRead"
)

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	corev1alpha1.ResourceClaimSpec `json:",inline"`

	// Name properties
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	Name string `json:"name,omitempty"`

	// +kubebuilder:validation:Enum=Private,PublicRead,PublicReadWrite,AuthenticatedRead
	// NOTE: AWS S3 and GCP Bucket values (not in Azure)
	PredefinedACL *PredefinedACL `json:"predefinedACL,omitempty"`

	// LocalPermission is the permissions granted on the bucket for the provider specific
	// bucket service account that is available in a secret after provisioning.
	// +kubebuilder:validation:Enum=Read,Write,ReadWrite
	// NOTE: AWS S3 Specific value
	LocalPermission *LocalPermissionType `json:"localPermission,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Bucket is the Schema for the Bucket API
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="PREDEFINED-ACL",type="string",JSONPath=".spec.predefinedACL"
// +kubebuilder:printcolumn:name="LOCAL-PERMISSION",type="string",JSONPath=".spec.localPermission"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec                       `json:"spec,omitempty"`
	Status corev1alpha1.ResourceClaimStatus `json:"status,omitempty"`
}

// SetBindingPhase of this Bucket.
func (b *Bucket) SetBindingPhase(p corev1alpha1.BindingPhase) {
	b.Status.SetBindingPhase(p)
}

// GetBindingPhase of this Bucket.
func (b *Bucket) GetBindingPhase() corev1alpha1.BindingPhase {
	return b.Status.GetBindingPhase()
}

// SetConditions of this Bucket.
func (b *Bucket) SetConditions(c ...corev1alpha1.Condition) {
	b.Status.SetConditions(c...)
}

// SetClassReference of this Bucket.
func (b *Bucket) SetClassReference(r *corev1.ObjectReference) {
	b.Spec.ClassReference = r
}

// GetClassReference of this Bucket.
func (b *Bucket) GetClassReference() *corev1.ObjectReference {
	return b.Spec.ClassReference
}

// SetResourceReference of this Bucket.
func (b *Bucket) SetResourceReference(r *corev1.ObjectReference) {
	b.Spec.ResourceReference = r
}

// GetResourceReference of this Bucket.
func (b *Bucket) GetResourceReference() *corev1.ObjectReference {
	return b.Spec.ResourceReference
}

// SetWriteConnectionSecretToReference of this Bucket.
func (b *Bucket) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	b.Spec.WriteConnectionSecretToReference = r
}

// GetWriteConnectionSecretToReference of this Bucket.
func (b *Bucket) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return b.Spec.WriteConnectionSecretToReference
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}
