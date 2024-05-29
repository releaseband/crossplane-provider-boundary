// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LibraryVaultSSHCertificateInitParameters struct {

	// (List of String) Principals to be signed as "valid_principles" in addition to username.
	// Principals to be signed as "valid_principles" in addition to username.
	AdditionalValidPrincipals []*string `json:"additionalValidPrincipals,omitempty" tf:"additional_valid_principals,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (Map of String) Specifies a map of the critical options that the certificate should be signed for.
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions map[string]*string `json:"criticalOptions,omitempty" tf:"critical_options,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Map of String) Specifies a map of the extensions that the certificate should be signed for.
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions map[string]*string `json:"extensions,omitempty" tf:"extensions,omitempty"`

	// (Number) Specifies the number of bits to use for the generated keys.
	// Specifies the number of bits to use for the generated keys.
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// (String) Specifies the key id a certificate should have.
	// Specifies the key id a certificate should have.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// (String) Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Specifies the requested time to live for a certificate returned from the library.
	// Specifies the requested time to live for a certificate returned from the library.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) The username to use with the certificate returned by the library.
	// The username to use with the certificate returned by the library.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LibraryVaultSSHCertificateObservation struct {

	// (List of String) Principals to be signed as "valid_principles" in addition to username.
	// Principals to be signed as "valid_principles" in addition to username.
	AdditionalValidPrincipals []*string `json:"additionalValidPrincipals,omitempty" tf:"additional_valid_principals,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (Map of String) Specifies a map of the critical options that the certificate should be signed for.
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions map[string]*string `json:"criticalOptions,omitempty" tf:"critical_options,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Map of String) Specifies a map of the extensions that the certificate should be signed for.
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions map[string]*string `json:"extensions,omitempty" tf:"extensions,omitempty"`

	// (String) The ID of the Vault credential library.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) Specifies the number of bits to use for the generated keys.
	// Specifies the number of bits to use for the generated keys.
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// (String) Specifies the key id a certificate should have.
	// Specifies the key id a certificate should have.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// (String) Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Specifies the requested time to live for a certificate returned from the library.
	// Specifies the requested time to live for a certificate returned from the library.
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) The username to use with the certificate returned by the library.
	// The username to use with the certificate returned by the library.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type LibraryVaultSSHCertificateParameters struct {

	// (List of String) Principals to be signed as "valid_principles" in addition to username.
	// Principals to be signed as "valid_principles" in addition to username.
	// +kubebuilder:validation:Optional
	AdditionalValidPrincipals []*string `json:"additionalValidPrincipals,omitempty" tf:"additional_valid_principals,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	// +kubebuilder:validation:Optional
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (Map of String) Specifies a map of the critical options that the certificate should be signed for.
	// Specifies a map of the critical options that the certificate should be signed for.
	// +kubebuilder:validation:Optional
	CriticalOptions map[string]*string `json:"criticalOptions,omitempty" tf:"critical_options,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Map of String) Specifies a map of the extensions that the certificate should be signed for.
	// Specifies a map of the extensions that the certificate should be signed for.
	// +kubebuilder:validation:Optional
	Extensions map[string]*string `json:"extensions,omitempty" tf:"extensions,omitempty"`

	// (Number) Specifies the number of bits to use for the generated keys.
	// Specifies the number of bits to use for the generated keys.
	// +kubebuilder:validation:Optional
	KeyBits *float64 `json:"keyBits,omitempty" tf:"key_bits,omitempty"`

	// (String) Specifies the key id a certificate should have.
	// Specifies the key id a certificate should have.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// (String) Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	// +kubebuilder:validation:Optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (String) Specifies the requested time to live for a certificate returned from the library.
	// Specifies the requested time to live for a certificate returned from the library.
	// +kubebuilder:validation:Optional
	TTL *string `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) The username to use with the certificate returned by the library.
	// The username to use with the certificate returned by the library.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// LibraryVaultSSHCertificateSpec defines the desired state of LibraryVaultSSHCertificate
type LibraryVaultSSHCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibraryVaultSSHCertificateParameters `json:"forProvider"`
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
	InitProvider LibraryVaultSSHCertificateInitParameters `json:"initProvider,omitempty"`
}

// LibraryVaultSSHCertificateStatus defines the observed state of LibraryVaultSSHCertificate.
type LibraryVaultSSHCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibraryVaultSSHCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryVaultSSHCertificate is the Schema for the LibraryVaultSSHCertificates API. The credential library for Vault resource allows you to configure a Boundary credential library for Vault.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type LibraryVaultSSHCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialStoreId) || (has(self.initProvider) && has(self.initProvider.credentialStoreId))",message="spec.forProvider.credentialStoreId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   LibraryVaultSSHCertificateSpec   `json:"spec"`
	Status LibraryVaultSSHCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryVaultSSHCertificateList contains a list of LibraryVaultSSHCertificates
type LibraryVaultSSHCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibraryVaultSSHCertificate `json:"items"`
}

// Repository type metadata.
var (
	LibraryVaultSSHCertificate_Kind             = "LibraryVaultSSHCertificate"
	LibraryVaultSSHCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LibraryVaultSSHCertificate_Kind}.String()
	LibraryVaultSSHCertificate_KindAPIVersion   = LibraryVaultSSHCertificate_Kind + "." + CRDGroupVersion.String()
	LibraryVaultSSHCertificate_GroupVersionKind = CRDGroupVersion.WithKind(LibraryVaultSSHCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&LibraryVaultSSHCertificate{}, &LibraryVaultSSHCertificateList{})
}
