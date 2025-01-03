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

type LibraryVaultInitParameters struct {

	// (Map of String) The credential mapping override.
	// The credential mapping override.
	// +mapType=granular
	CredentialMappingOverrides map[string]*string `json:"credentialMappingOverrides,omitempty" tf:"credential_mapping_overrides,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The type of credential the library generates. Cannot be updated on an existing resource.
	// The type of credential the library generates. Cannot be updated on an existing resource.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	// The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (String) The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if http_method is set to POST.
	// The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `http_method` is set to `POST`.
	HTTPRequestBody *string `json:"httpRequestBody,omitempty" tf:"http_request_body,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LibraryVaultObservation struct {

	// (Map of String) The credential mapping override.
	// The credential mapping override.
	// +mapType=granular
	CredentialMappingOverrides map[string]*string `json:"credentialMappingOverrides,omitempty" tf:"credential_mapping_overrides,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The type of credential the library generates. Cannot be updated on an existing resource.
	// The type of credential the library generates. Cannot be updated on an existing resource.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	// The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (String) The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if http_method is set to POST.
	// The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `http_method` is set to `POST`.
	HTTPRequestBody *string `json:"httpRequestBody,omitempty" tf:"http_request_body,omitempty"`

	// (String) The ID of the Vault credential library.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type LibraryVaultParameters struct {

	// (Map of String) The credential mapping override.
	// The credential mapping override.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	CredentialMappingOverrides map[string]*string `json:"credentialMappingOverrides,omitempty" tf:"credential_mapping_overrides,omitempty"`

	// (String) The ID of the credential store that this library belongs to.
	// The ID of the credential store that this library belongs to.
	// +kubebuilder:validation:Optional
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The type of credential the library generates. Cannot be updated on an existing resource.
	// The type of credential the library generates. Cannot be updated on an existing resource.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String) The Vault credential library description.
	// The Vault credential library description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	// The HTTP method the library uses when requesting credentials from Vault. Defaults to 'GET'
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// (String) The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if http_method is set to POST.
	// The body of the HTTP request the library sends to Vault when requesting credentials. Only valid if `http_method` is set to `POST`.
	// +kubebuilder:validation:Optional
	HTTPRequestBody *string `json:"httpRequestBody,omitempty" tf:"http_request_body,omitempty"`

	// (String) The Vault credential library name. Defaults to the resource name.
	// The Vault credential library name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The path in Vault to request credentials from.
	// The path in Vault to request credentials from.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

// LibraryVaultSpec defines the desired state of LibraryVault
type LibraryVaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibraryVaultParameters `json:"forProvider"`
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
	InitProvider LibraryVaultInitParameters `json:"initProvider,omitempty"`
}

// LibraryVaultStatus defines the observed state of LibraryVault.
type LibraryVaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibraryVaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LibraryVault is the Schema for the LibraryVaults API. The credential library for Vault resource allows you to configure a Boundary credential library for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type LibraryVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialStoreId) || (has(self.initProvider) && has(self.initProvider.credentialStoreId))",message="spec.forProvider.credentialStoreId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	Spec   LibraryVaultSpec   `json:"spec"`
	Status LibraryVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryVaultList contains a list of LibraryVaults
type LibraryVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibraryVault `json:"items"`
}

// Repository type metadata.
var (
	LibraryVault_Kind             = "LibraryVault"
	LibraryVault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LibraryVault_Kind}.String()
	LibraryVault_KindAPIVersion   = LibraryVault_Kind + "." + CRDGroupVersion.String()
	LibraryVault_GroupVersionKind = CRDGroupVersion.WithKind(LibraryVault_Kind)
)

func init() {
	SchemeBuilder.Register(&LibraryVault{}, &LibraryVaultList{})
}
