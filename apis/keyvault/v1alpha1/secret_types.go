/*
Copyright 2019 The Crossplane Authors.

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

const (
	// SupportedSecretVersion is the only minor version of Secret currently
	// supported by Azure Cache for Secret. The version cannot be specified at
	// creation time.
	SupportedSecretVersion = "3.2"
)

// DeletionRecoveryLevel enumerates the values for deletion recovery level.
type DeletionRecoveryLevel string

const (
	// CustomizedRecoverable Denotes a vault state in which deletion is recoverable without the possibility for
	// immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90).This level
	// guarantees the recoverability of the deleted entity during the retention interval and while the
	// subscription is still available.
	CustomizedRecoverable DeletionRecoveryLevel = "CustomizedRecoverable"

	// CustomizedRecoverableProtectedSubscription Denotes a vault and subscription state in which deletion is
	// recoverable, immediate and permanent deletion (i.e. purge) is not permitted, and in which the
	// subscription itself cannot be permanently canceled when 7<= SoftDeleteRetentionInDays < 90. This level
	// guarantees the recoverability of the deleted entity during the retention interval, and also reflects the
	// fact that the subscription itself cannot be cancelled.
	CustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = "CustomizedRecoverable+ProtectedSubscription"

	// CustomizedRecoverablePurgeable Denotes a vault state in which deletion is recoverable, and which also
	// permits immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90). This
	// level guarantees the recoverability of the deleted entity during the retention interval, unless a Purge
	// operation is requested, or the subscription is cancelled.
	CustomizedRecoverablePurgeable DeletionRecoveryLevel = "CustomizedRecoverable+Purgeable"

	// Purgeable Denotes a vault state in which deletion is an irreversible operation, without the possibility
	// for recovery. This level corresponds to no protection being available against a Delete operation; the
	// data is irretrievably lost upon accepting a Delete operation at the entity level or higher (vault,
	// resource group, subscription etc.)
	Purgeable DeletionRecoveryLevel = "Purgeable"

	// Recoverable Denotes a vault state in which deletion is recoverable without the possibility for immediate
	// and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted entity
	// during the retention interval(90 days) and while the subscription is still available. System wil
	// permanently delete it after 90 days, if not recovered
	Recoverable DeletionRecoveryLevel = "Recoverable"

	// RecoverableProtectedSubscription Denotes a vault and subscription state in which deletion is recoverable
	// within retention interval (90 days), immediate and permanent deletion (i.e. purge) is not permitted, and
	// in which the subscription itself  cannot be permanently canceled. System wil permanently delete it after
	// 90 days, if not recovered
	RecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"

	// RecoverablePurgeable Denotes a vault state in which deletion is recoverable, and which also permits
	// immediate and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted
	// entity during the retention interval (90 days), unless a Purge operation is requested, or the
	// subscription is cancelled. System wil permanently delete it after 90 days, if not recovered
	RecoverablePurgeable DeletionRecoveryLevel = "Recoverable+Purgeable"
)

type SecretAttributes struct {
	// RecoverableDays - READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32 `json:"recoverableDays,omitempty"`

	// RecoveryLevel - READ-ONLY; Reflects the deletion recovery level currently in effect for secrets in the current vault. If it contains 'Purgeable', the secret can be permanently deleted by a privileged user; otherwise, only the system can purge the secret, at the end of the retention interval. Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription', 'CustomizedRecoverablePurgeable', 'CustomizedRecoverable', 'CustomizedRecoverableProtectedSubscription'
	RecoveryLevel DeletionRecoveryLevel `json:"recoveryLevel,omitempty"`

	// Enabled - Determines whether the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// NotBefore - Not before date in UTC.
	NotBefore *metav1.Time `json:"nbf,omitempty"`

	// Expires - Expiry date in UTC.
	Expires *metav1.Time `json:"exp,omitempty"`

	// Created - READ-ONLY; Creation time in UTC.
	Created *metav1.Time `json:"created,omitempty"`

	// Updated - READ-ONLY; Last updated time in UTC.
	Updated *metav1.Time `json:"updated,omitempty"`
}

// SecretParameters define the desired state of an Azure Key Vault Secret.
// https://docs.microsoft.com/en-us/rest/api/keyvault/#secret-operations
type SecretParameters struct {
	// Name - The name of the secret
	Name string `json:"name"`

	// Value - The value of the secret
	Value string `json:"value"`

	// ContentType - Type of the secret value such as a password
	// +optional
	ContentType *string `json:"contentType,omitempty"`

	// SecretAttributes - The secret management attributes
	// +optional
	SecretAttributes *SecretAttributes `json:"attributes,omitempty"`

	// Tags - Application-specific metadata in the form of key-value pairs
	// +optional
	Tags map[string]string `json:"tags"`
}

// A SecretSpec defines the desired state of a Secret.
type SecretSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecretParameters `json:"forProvider"`
}

// SecretObservation represents the observed state of the Secret object in Azure.
type SecretObservation struct {
	// ID - The secret id.
	ID string `json:"id,omitempty"`

	// Value - The secret value.
	Value string `json:"value,omitempty"`

	// ContentType - The content type of the secret.
	ContentType string `json:"contentType,omitempty"`

	// Attributes - The secret management attributes.
	Attributes SecretAttributes `json:"attributes,omitempty"`

	// Kid - READ-ONLY; If this is a secret backing a KV certificate, then this field specifies the corresponding key backing the KV certificate.
	Kid string `json:"kid,omitempty"`

	// Managed - READ-ONLY; True if the secret's lifetime is managed by key vault. If this is a secret backing a certificate, then managed will be true.
	Managed bool `json:"managed,omitempty"`

	// Tags - Application specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags"`
}

// A SecretStatus represents the observed state of a Secret.
type SecretStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Secret is a managed resource that represents an Azure Secret cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretSpec   `json:"spec"`
	Status SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secret.
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}