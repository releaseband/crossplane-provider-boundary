// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SSHPrivateKeyInitParameters struct {

	// (String) ID of the credential store this credential belongs to.
	// ID of the credential store this credential belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of the credential.
	// The description of the credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of the credential. Defaults to the resource name.
	// The name of the credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The passphrase of the private key associated with the credential.
	// The passphrase of the private key associated with the credential.
	PrivateKeyPassphraseSecretRef *v1.SecretKeySelector `json:"privateKeyPassphraseSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) The private key associated with the credential.
	// The private key associated with the credential.
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (String) The username associated with the credential.
	// The username associated with the credential.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SSHPrivateKeyObservation struct {

	// (String) ID of the credential store this credential belongs to.
	// ID of the credential store this credential belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of the credential.
	// The description of the credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the credential.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the credential. Defaults to the resource name.
	// The name of the credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The private key hmac.
	// The private key hmac.
	PrivateKeyHMAC *string `json:"privateKeyHmac,omitempty" tf:"private_key_hmac,omitempty"`

	// (String) The private key passphrase hmac.
	// The private key passphrase hmac.
	PrivateKeyPassphraseHMAC *string `json:"privateKeyPassphraseHmac,omitempty" tf:"private_key_passphrase_hmac,omitempty"`

	// (String) The username associated with the credential.
	// The username associated with the credential.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SSHPrivateKeyParameters struct {

	// (String) ID of the credential store this credential belongs to.
	// ID of the credential store this credential belongs to.
	// +kubebuilder:validation:Optional
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of the credential.
	// The description of the credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of the credential. Defaults to the resource name.
	// The name of the credential. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The passphrase of the private key associated with the credential.
	// The passphrase of the private key associated with the credential.
	// +kubebuilder:validation:Optional
	PrivateKeyPassphraseSecretRef *v1.SecretKeySelector `json:"privateKeyPassphraseSecretRef,omitempty" tf:"-"`

	// (String, Sensitive) The private key associated with the credential.
	// The private key associated with the credential.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (String) The username associated with the credential.
	// The username associated with the credential.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// SSHPrivateKeySpec defines the desired state of SSHPrivateKey
type SSHPrivateKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHPrivateKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SSHPrivateKeyInitParameters `json:"initProvider,omitempty"`
}

// SSHPrivateKeyStatus defines the observed state of SSHPrivateKey.
type SSHPrivateKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHPrivateKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSHPrivateKey is the Schema for the SSHPrivateKeys API. The SSH private key credential resource allows you to configure a credential using a username, private key and optional passphrase.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type SSHPrivateKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialStoreId) || (has(self.initProvider) && has(self.initProvider.credentialStoreId))",message="spec.forProvider.credentialStoreId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   SSHPrivateKeySpec   `json:"spec"`
	Status SSHPrivateKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHPrivateKeyList contains a list of SSHPrivateKeys
type SSHPrivateKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHPrivateKey `json:"items"`
}

// Repository type metadata.
var (
	SSHPrivateKey_Kind             = "SSHPrivateKey"
	SSHPrivateKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHPrivateKey_Kind}.String()
	SSHPrivateKey_KindAPIVersion   = SSHPrivateKey_Kind + "." + CRDGroupVersion.String()
	SSHPrivateKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHPrivateKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHPrivateKey{}, &SSHPrivateKeyList{})
}
