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

type UsernamePasswordInitParameters struct {

	// (String) The credential store in which to save this username/password credential.
	// The credential store in which to save this username/password credential.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this username/password credential.
	// The description of this username/password credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of this username/password credential. Defaults to the resource name.
	// The name of this username/password credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The password of this username/password credential.
	// The password of this username/password credential.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// (String) The username of this username/password credential.
	// The username of this username/password credential.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UsernamePasswordObservation struct {

	// (String) The credential store in which to save this username/password credential.
	// The credential store in which to save this username/password credential.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this username/password credential.
	// The description of this username/password credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this username/password credential.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of this username/password credential. Defaults to the resource name.
	// The name of this username/password credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The password hmac.
	// The password hmac.
	PasswordHMAC *string `json:"passwordHmac,omitempty" tf:"password_hmac,omitempty"`

	// (String) The username of this username/password credential.
	// The username of this username/password credential.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UsernamePasswordParameters struct {

	// (String) The credential store in which to save this username/password credential.
	// The credential store in which to save this username/password credential.
	// +kubebuilder:validation:Optional
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this username/password credential.
	// The description of this username/password credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of this username/password credential. Defaults to the resource name.
	// The name of this username/password credential. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The password of this username/password credential.
	// The password of this username/password credential.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// (String) The username of this username/password credential.
	// The username of this username/password credential.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UsernamePasswordSpec defines the desired state of UsernamePassword
type UsernamePasswordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UsernamePasswordParameters `json:"forProvider"`
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
	InitProvider UsernamePasswordInitParameters `json:"initProvider,omitempty"`
}

// UsernamePasswordStatus defines the observed state of UsernamePassword.
type UsernamePasswordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UsernamePasswordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UsernamePassword is the Schema for the UsernamePasswords API. The username/password credential resource allows you to configure a credential using a username and password pair.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type UsernamePassword struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialStoreId) || (has(self.initProvider) && has(self.initProvider.credentialStoreId))",message="spec.forProvider.credentialStoreId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   UsernamePasswordSpec   `json:"spec"`
	Status UsernamePasswordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UsernamePasswordList contains a list of UsernamePasswords
type UsernamePasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UsernamePassword `json:"items"`
}

// Repository type metadata.
var (
	UsernamePassword_Kind             = "UsernamePassword"
	UsernamePassword_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UsernamePassword_Kind}.String()
	UsernamePassword_KindAPIVersion   = UsernamePassword_Kind + "." + CRDGroupVersion.String()
	UsernamePassword_GroupVersionKind = CRDGroupVersion.WithKind(UsernamePassword_Kind)
)

func init() {
	SchemeBuilder.Register(&UsernamePassword{}, &UsernamePasswordList{})
}
