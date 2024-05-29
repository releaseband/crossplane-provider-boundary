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

type SetInitParameters struct {

	// (String) The host set description.
	// The host set description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The catalog for the host set.
	// The catalog for the host set.
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (Set of String) The list of host IDs contained in this set.
	// The list of host IDs contained in this set.
	HostIds []*string `json:"hostIds,omitempty" tf:"host_ids,omitempty"`

	// (String) The host set name. Defaults to the resource name.
	// The host set name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host set
	// The type of host set
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SetObservation struct {

	// (String) The host set description.
	// The host set description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The catalog for the host set.
	// The catalog for the host set.
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (Set of String) The list of host IDs contained in this set.
	// The list of host IDs contained in this set.
	HostIds []*string `json:"hostIds,omitempty" tf:"host_ids,omitempty"`

	// (String) The ID of the host set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The host set name. Defaults to the resource name.
	// The host set name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host set
	// The type of host set
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SetParameters struct {

	// (String) The host set description.
	// The host set description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The catalog for the host set.
	// The catalog for the host set.
	// +kubebuilder:validation:Optional
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (Set of String) The list of host IDs contained in this set.
	// The list of host IDs contained in this set.
	// +kubebuilder:validation:Optional
	HostIds []*string `json:"hostIds,omitempty" tf:"host_ids,omitempty"`

	// (String) The host set name. Defaults to the resource name.
	// The host set name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host set
	// The type of host set
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SetSpec defines the desired state of Set
type SetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SetParameters `json:"forProvider"`
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
	InitProvider SetInitParameters `json:"initProvider,omitempty"`
}

// SetStatus defines the observed state of Set.
type SetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Set is the Schema for the Sets API. Deprecated: use boundary_host_set_static instead.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Set struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostCatalogId) || (has(self.initProvider) && has(self.initProvider.hostCatalogId))",message="spec.forProvider.hostCatalogId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SetSpec   `json:"spec"`
	Status SetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SetList contains a list of Sets
type SetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Set `json:"items"`
}

// Repository type metadata.
var (
	Set_Kind             = "Set"
	Set_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Set_Kind}.String()
	Set_KindAPIVersion   = Set_Kind + "." + CRDGroupVersion.String()
	Set_GroupVersionKind = CRDGroupVersion.WithKind(Set_Kind)
)

func init() {
	SchemeBuilder.Register(&Set{}, &SetList{})
}