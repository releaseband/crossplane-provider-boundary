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

type StorageInitParameters struct {

	// The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically delete. However, delete_after_days and retain_for_days cannot both be 0.
	DeleteAfterDays *float64 `json:"deleteAfterDays,omitempty" tf:"delete_after_days,omitempty"`

	// Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated delete_after_days value is 0, overridable is ignored
	DeleteAfterOverridable *bool `json:"deleteAfterOverridable,omitempty" tf:"delete_after_overridable,omitempty"`

	// The policy description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The policy name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However, retain_for_days and delete_after_days cannot both be 0.
	RetainForDays *float64 `json:"retainForDays,omitempty" tf:"retain_for_days,omitempty"`

	// Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated retain_for_days value is 0, overridable is ignored.
	RetainForOverridable *bool `json:"retainForOverridable,omitempty" tf:"retain_for_overridable,omitempty"`

	// The scope for this policy.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type StorageObservation struct {

	// The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically delete. However, delete_after_days and retain_for_days cannot both be 0.
	DeleteAfterDays *float64 `json:"deleteAfterDays,omitempty" tf:"delete_after_days,omitempty"`

	// Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated delete_after_days value is 0, overridable is ignored
	DeleteAfterOverridable *bool `json:"deleteAfterOverridable,omitempty" tf:"delete_after_overridable,omitempty"`

	// The policy description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However, retain_for_days and delete_after_days cannot both be 0.
	RetainForDays *float64 `json:"retainForDays,omitempty" tf:"retain_for_days,omitempty"`

	// Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated retain_for_days value is 0, overridable is ignored.
	RetainForOverridable *bool `json:"retainForOverridable,omitempty" tf:"retain_for_overridable,omitempty"`

	// The scope for this policy.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type StorageParameters struct {

	// The number of days after which a session recording will be automatically deleted. Defaults to 0: never automatically delete. However, delete_after_days and retain_for_days cannot both be 0.
	// +kubebuilder:validation:Optional
	DeleteAfterDays *float64 `json:"deleteAfterDays,omitempty" tf:"delete_after_days,omitempty"`

	// Whether or not the associated delete_after_days value can be overridden by org scopes. Note: if the associated delete_after_days value is 0, overridable is ignored
	// +kubebuilder:validation:Optional
	DeleteAfterOverridable *bool `json:"deleteAfterOverridable,omitempty" tf:"delete_after_overridable,omitempty"`

	// The policy description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The policy name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of days a session recording is required to be stored. Defaults to 0: allow deletions at any time. However, retain_for_days and delete_after_days cannot both be 0.
	// +kubebuilder:validation:Optional
	RetainForDays *float64 `json:"retainForDays,omitempty" tf:"retain_for_days,omitempty"`

	// Whether or not the associated retain_for_days value can be overridden by org scopes. Note: if the associated retain_for_days value is 0, overridable is ignored.
	// +kubebuilder:validation:Optional
	RetainForOverridable *bool `json:"retainForOverridable,omitempty" tf:"retain_for_overridable,omitempty"`

	// The scope for this policy.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

// StorageSpec defines the desired state of Storage
type StorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageParameters `json:"forProvider"`
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
	InitProvider StorageInitParameters `json:"initProvider,omitempty"`
}

// StorageStatus defines the observed state of Storage.
type StorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Storage is the Schema for the Storages API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	Spec   StorageSpec   `json:"spec"`
	Status StorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageList contains a list of Storages
type StorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Storage `json:"items"`
}

// Repository type metadata.
var (
	Storage_Kind             = "Storage"
	Storage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Storage_Kind}.String()
	Storage_KindAPIVersion   = Storage_Kind + "." + CRDGroupVersion.String()
	Storage_GroupVersionKind = CRDGroupVersion.WithKind(Storage_Kind)
)

func init() {
	SchemeBuilder.Register(&Storage{}, &StorageList{})
}
