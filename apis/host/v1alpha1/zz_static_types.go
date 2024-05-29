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

type StaticInitParameters struct {

	// (String) The static address of the host resource as <IP> (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (String) The host description.
	// The host description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (String) The host name. Defaults to the resource name.
	// The host name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host
	// The type of host
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StaticObservation struct {

	// (String) The static address of the host resource as <IP> (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (String) The host description.
	// The host description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (String) The ID of the host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The host name. Defaults to the resource name.
	// The host name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host
	// The type of host
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StaticParameters struct {

	// (String) The static address of the host resource as <IP> (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	// The static address of the host resource as `<IP>` (note: port assignment occurs in the target resource definition, do not add :port here) or a domain name.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (String) The host description.
	// The host description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	HostCatalogID *string `json:"hostCatalogId,omitempty" tf:"host_catalog_id,omitempty"`

	// (String) The host name. Defaults to the resource name.
	// The host name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of host
	// The type of host
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// StaticSpec defines the desired state of Static
type StaticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StaticParameters `json:"forProvider"`
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
	InitProvider StaticInitParameters `json:"initProvider,omitempty"`
}

// StaticStatus defines the observed state of Static.
type StaticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StaticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Static is the Schema for the Statics API. The static host resource allows you to configure a Boundary static host. Hosts are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Static struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostCatalogId) || (has(self.initProvider) && has(self.initProvider.hostCatalogId))",message="spec.forProvider.hostCatalogId is a required parameter"
	Spec   StaticSpec   `json:"spec"`
	Status StaticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticList contains a list of Statics
type StaticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Static `json:"items"`
}

// Repository type metadata.
var (
	Static_Kind             = "Static"
	Static_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Static_Kind}.String()
	Static_KindAPIVersion   = Static_Kind + "." + CRDGroupVersion.String()
	Static_GroupVersionKind = CRDGroupVersion.WithKind(Static_Kind)
)

func init() {
	SchemeBuilder.Register(&Static{}, &StaticList{})
}