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

type StoreStaticInitParameters struct {

	// (String) The static credential store description.
	// The static credential store description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The static credential store name. Defaults to the resource name.
	// The static credential store name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope for this credential store.
	// The scope for this credential store.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type StoreStaticObservation struct {

	// (String) The static credential store description.
	// The static credential store description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the static credential store.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The static credential store name. Defaults to the resource name.
	// The static credential store name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope for this credential store.
	// The scope for this credential store.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type StoreStaticParameters struct {

	// (String) The static credential store description.
	// The static credential store description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The static credential store name. Defaults to the resource name.
	// The static credential store name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope for this credential store.
	// The scope for this credential store.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

// StoreStaticSpec defines the desired state of StoreStatic
type StoreStaticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StoreStaticParameters `json:"forProvider"`
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
	InitProvider StoreStaticInitParameters `json:"initProvider,omitempty"`
}

// StoreStaticStatus defines the observed state of StoreStatic.
type StoreStaticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StoreStaticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StoreStatic is the Schema for the StoreStatics API. The static credential store resource allows you to configure a Boundary static credential store.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type StoreStatic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	Spec   StoreStaticSpec   `json:"spec"`
	Status StoreStaticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoreStaticList contains a list of StoreStatics
type StoreStaticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoreStatic `json:"items"`
}

// Repository type metadata.
var (
	StoreStatic_Kind             = "StoreStatic"
	StoreStatic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StoreStatic_Kind}.String()
	StoreStatic_KindAPIVersion   = StoreStatic_Kind + "." + CRDGroupVersion.String()
	StoreStatic_GroupVersionKind = CRDGroupVersion.WithKind(StoreStatic_Kind)
)

func init() {
	SchemeBuilder.Register(&StoreStatic{}, &StoreStaticList{})
}
