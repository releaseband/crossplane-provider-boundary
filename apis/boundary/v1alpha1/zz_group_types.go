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

type GroupInitParameters struct {

	// (String) The group description.
	// The group description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) Resource IDs for group members, these are most likely boundary users.
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// (String) The group name. Defaults to the resource name.
	// The group name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type GroupObservation struct {

	// (String) The group description.
	// The group description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) Resource IDs for group members, these are most likely boundary users.
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// (String) The group name. Defaults to the resource name.
	// The group name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type GroupParameters struct {

	// (String) The group description.
	// The group description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) Resource IDs for group members, these are most likely boundary users.
	// Resource IDs for group members, these are most likely boundary users.
	// +kubebuilder:validation:Optional
	MemberIds []*string `json:"memberIds,omitempty" tf:"member_ids,omitempty"`

	// (String) The group name. Defaults to the resource name.
	// The group name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. The group resource allows you to configure a Boundary group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
