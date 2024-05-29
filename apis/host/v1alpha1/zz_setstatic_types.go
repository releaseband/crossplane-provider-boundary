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

type SetStaticInitParameters struct {

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

type SetStaticObservation struct {

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

type SetStaticParameters struct {

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

// SetStaticSpec defines the desired state of SetStatic
type SetStaticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SetStaticParameters `json:"forProvider"`
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
	InitProvider SetStaticInitParameters `json:"initProvider,omitempty"`
}

// SetStaticStatus defines the observed state of SetStatic.
type SetStaticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SetStaticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SetStatic is the Schema for the SetStatics API. The hostsetstatic resource allows you to configure a Boundary host set. Host sets are always part of a host catalog, so a host catalog resource should be used inline or you should have the host catalog ID in hand to successfully configure a host set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type SetStatic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostCatalogId) || (has(self.initProvider) && has(self.initProvider.hostCatalogId))",message="spec.forProvider.hostCatalogId is a required parameter"
	Spec   SetStaticSpec   `json:"spec"`
	Status SetStaticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SetStaticList contains a list of SetStatics
type SetStaticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SetStatic `json:"items"`
}

// Repository type metadata.
var (
	SetStatic_Kind             = "SetStatic"
	SetStatic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SetStatic_Kind}.String()
	SetStatic_KindAPIVersion   = SetStatic_Kind + "." + CRDGroupVersion.String()
	SetStatic_GroupVersionKind = CRDGroupVersion.WithKind(SetStatic_Kind)
)

func init() {
	SchemeBuilder.Register(&SetStatic{}, &SetStaticList{})
}