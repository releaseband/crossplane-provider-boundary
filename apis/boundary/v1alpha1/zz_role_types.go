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

type RoleInitParameters struct {

	// (String) The role description.
	// The role description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String, Deprecated) For Boundary 0.15+, use grant_scope_ids instead. The scope for which the grants in the role should apply.
	// For Boundary 0.15+, use `grant_scope_ids` instead. The scope for which the grants in the role should apply.
	GrantScopeID *string `json:"grantScopeId,omitempty" tf:"grant_scope_id,omitempty"`

	// (Set of String) A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds []*string `json:"grantScopeIds,omitempty" tf:"grant_scope_ids,omitempty"`

	// (Set of String) A list of stringified grants for the role.
	// A list of stringified grants for the role.
	GrantStrings []*string `json:"grantStrings,omitempty" tf:"grant_strings,omitempty"`

	// (String) The role name. Defaults to the resource name.
	// The role name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of principal (user or group) IDs to add as principals on the role.
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []*string `json:"principalIds,omitempty" tf:"principal_ids,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type RoleObservation struct {

	// (String) The role description.
	// The role description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String, Deprecated) For Boundary 0.15+, use grant_scope_ids instead. The scope for which the grants in the role should apply.
	// For Boundary 0.15+, use `grant_scope_ids` instead. The scope for which the grants in the role should apply.
	GrantScopeID *string `json:"grantScopeId,omitempty" tf:"grant_scope_id,omitempty"`

	// (Set of String) A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds []*string `json:"grantScopeIds,omitempty" tf:"grant_scope_ids,omitempty"`

	// (Set of String) A list of stringified grants for the role.
	// A list of stringified grants for the role.
	GrantStrings []*string `json:"grantStrings,omitempty" tf:"grant_strings,omitempty"`

	// (String) The ID of the role.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The role name. Defaults to the resource name.
	// The role name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of principal (user or group) IDs to add as principals on the role.
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []*string `json:"principalIds,omitempty" tf:"principal_ids,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type RoleParameters struct {

	// (String) The role description.
	// The role description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String, Deprecated) For Boundary 0.15+, use grant_scope_ids instead. The scope for which the grants in the role should apply.
	// For Boundary 0.15+, use `grant_scope_ids` instead. The scope for which the grants in the role should apply.
	// +kubebuilder:validation:Optional
	GrantScopeID *string `json:"grantScopeId,omitempty" tf:"grant_scope_id,omitempty"`

	// (Set of String) A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	// +kubebuilder:validation:Optional
	GrantScopeIds []*string `json:"grantScopeIds,omitempty" tf:"grant_scope_ids,omitempty"`

	// (Set of String) A list of stringified grants for the role.
	// A list of stringified grants for the role.
	// +kubebuilder:validation:Optional
	GrantStrings []*string `json:"grantStrings,omitempty" tf:"grant_strings,omitempty"`

	// (String) The role name. Defaults to the resource name.
	// The role name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of principal (user or group) IDs to add as principals on the role.
	// A list of principal (user or group) IDs to add as principals on the role.
	// +kubebuilder:validation:Optional
	PrincipalIds []*string `json:"principalIds,omitempty" tf:"principal_ids,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Role is the Schema for the Roles API. The role resource allows you to configure a Boundary role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}