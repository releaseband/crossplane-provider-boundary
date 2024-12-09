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

type CatalogStaticInitParameters struct {

	// (String) The host catalog description.
	// The host catalog description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	// +crossplane:generate:reference:type=github.com/releaseband/crossplane-provider-boundary/apis/main/v1alpha1.Scope
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// Reference to a Scope in main to populate scopeId.
	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// Selector for a Scope in main to populate scopeId.
	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`
}

type CatalogStaticObservation struct {

	// (String) The host catalog description.
	// The host catalog description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the host catalog.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type CatalogStaticParameters struct {

	// (String) The host catalog description.
	// The host catalog description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	// +crossplane:generate:reference:type=github.com/releaseband/crossplane-provider-boundary/apis/main/v1alpha1.Scope
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// Reference to a Scope in main to populate scopeId.
	// +kubebuilder:validation:Optional
	ScopeIDRef *v1.Reference `json:"scopeIdRef,omitempty" tf:"-"`

	// Selector for a Scope in main to populate scopeId.
	// +kubebuilder:validation:Optional
	ScopeIDSelector *v1.Selector `json:"scopeIdSelector,omitempty" tf:"-"`
}

// CatalogStaticSpec defines the desired state of CatalogStatic
type CatalogStaticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogStaticParameters `json:"forProvider"`
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
	InitProvider CatalogStaticInitParameters `json:"initProvider,omitempty"`
}

// CatalogStaticStatus defines the observed state of CatalogStatic.
type CatalogStaticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogStaticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CatalogStatic is the Schema for the CatalogStatics API. The static host catalog resource allows you to configure a Boundary static-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type CatalogStatic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogStaticSpec   `json:"spec"`
	Status            CatalogStaticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogStaticList contains a list of CatalogStatics
type CatalogStaticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogStatic `json:"items"`
}

// Repository type metadata.
var (
	CatalogStatic_Kind             = "CatalogStatic"
	CatalogStatic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogStatic_Kind}.String()
	CatalogStatic_KindAPIVersion   = CatalogStatic_Kind + "." + CRDGroupVersion.String()
	CatalogStatic_GroupVersionKind = CRDGroupVersion.WithKind(CatalogStatic_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogStatic{}, &CatalogStaticList{})
}
