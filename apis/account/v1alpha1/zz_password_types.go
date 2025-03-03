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

type PasswordInitParameters struct {

	// (String) The resource ID for the auth method.
	// The resource ID for the auth method.
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// (String) The account description.
	// The account description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The login name for this account.
	// The login name for this account.
	LoginName *string `json:"loginName,omitempty" tf:"login_name,omitempty"`

	// (String) The account name. Defaults to the resource name.
	// The account name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The account password. Only set on create, changes will not be reflected when updating account.
	// The account password. Only set on create, changes will not be reflected when updating account.
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String, Deprecated) The resource type.
	// The resource type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PasswordObservation struct {

	// (String) The resource ID for the auth method.
	// The resource ID for the auth method.
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// (String) The account description.
	// The account description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The login name for this account.
	// The login name for this account.
	LoginName *string `json:"loginName,omitempty" tf:"login_name,omitempty"`

	// (String) The account name. Defaults to the resource name.
	// The account name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Deprecated) The resource type.
	// The resource type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PasswordParameters struct {

	// (String) The resource ID for the auth method.
	// The resource ID for the auth method.
	// +kubebuilder:validation:Optional
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// (String) The account description.
	// The account description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The login name for this account.
	// The login name for this account.
	// +kubebuilder:validation:Optional
	LoginName *string `json:"loginName,omitempty" tf:"login_name,omitempty"`

	// (String) The account name. Defaults to the resource name.
	// The account name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The account password. Only set on create, changes will not be reflected when updating account.
	// The account password. Only set on create, changes will not be reflected when updating account.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String, Deprecated) The resource type.
	// The resource type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// PasswordSpec defines the desired state of Password
type PasswordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PasswordParameters `json:"forProvider"`
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
	InitProvider PasswordInitParameters `json:"initProvider,omitempty"`
}

// PasswordStatus defines the observed state of Password.
type PasswordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PasswordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Password is the Schema for the Passwords API. The account resource allows you to configure a Boundary account.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Password struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authMethodId) || (has(self.initProvider) && has(self.initProvider.authMethodId))",message="spec.forProvider.authMethodId is a required parameter"
	Spec   PasswordSpec   `json:"spec"`
	Status PasswordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordList contains a list of Passwords
type PasswordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Password `json:"items"`
}

// Repository type metadata.
var (
	Password_Kind             = "Password"
	Password_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Password_Kind}.String()
	Password_KindAPIVersion   = Password_Kind + "." + CRDGroupVersion.String()
	Password_GroupVersionKind = CRDGroupVersion.WithKind(Password_Kind)
)

func init() {
	SchemeBuilder.Register(&Password{}, &PasswordList{})
}
