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

type TargetInitParameters struct {

	// (String) Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	// Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (Set of String) A list of brokered credential source ID's.
	// A list of brokered credential source ID's.
	BrokeredCredentialSourceIds []*string `json:"brokeredCredentialSourceIds,omitempty" tf:"brokered_credential_source_ids,omitempty"`

	// (Number) The default client port for this target.
	// The default client port for this target.
	DefaultClientPort *float64 `json:"defaultClientPort,omitempty" tf:"default_client_port,omitempty"`

	// (Number) The default port for this target.
	// The default port for this target.
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// (String) The target description.
	// The target description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Boolean expression to filter the workers used to access this target
	// Boolean expression to filter the workers used to access this target
	EgressWorkerFilter *string `json:"egressWorkerFilter,omitempty" tf:"egress_worker_filter,omitempty"`

	// (Boolean) HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	EnableSessionRecording *bool `json:"enableSessionRecording,omitempty" tf:"enable_session_recording,omitempty"`

	// (Set of String) A list of host source ID's. Cannot be used alongside address.
	// A list of host source ID's. Cannot be used alongside address.
	HostSourceIds []*string `json:"hostSourceIds,omitempty" tf:"host_source_ids,omitempty"`

	// (String) HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	// HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	IngressWorkerFilter *string `json:"ingressWorkerFilter,omitempty" tf:"ingress_worker_filter,omitempty"`

	// (Set of String) A list of injected application credential source ID's.
	// A list of injected application credential source ID's.
	InjectedApplicationCredentialSourceIds []*string `json:"injectedApplicationCredentialSourceIds,omitempty" tf:"injected_application_credential_source_ids,omitempty"`

	// (String) The target name. Defaults to the resource name.
	// The target name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (Number)
	SessionConnectionLimit *float64 `json:"sessionConnectionLimit,omitempty" tf:"session_connection_limit,omitempty"`

	// (Number)
	SessionMaxSeconds *float64 `json:"sessionMaxSeconds,omitempty" tf:"session_max_seconds,omitempty"`

	// (String) HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	StorageBucketID *string `json:"storageBucketId,omitempty" tf:"storage_bucket_id,omitempty"`

	// (String) The target resource type.
	// The target resource type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String, Deprecated) Boolean expression to filter the workers for this target
	// Boolean expression to filter the workers for this target
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

type TargetObservation struct {

	// (String) Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	// Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (Set of String) A list of brokered credential source ID's.
	// A list of brokered credential source ID's.
	BrokeredCredentialSourceIds []*string `json:"brokeredCredentialSourceIds,omitempty" tf:"brokered_credential_source_ids,omitempty"`

	// (Number) The default client port for this target.
	// The default client port for this target.
	DefaultClientPort *float64 `json:"defaultClientPort,omitempty" tf:"default_client_port,omitempty"`

	// (Number) The default port for this target.
	// The default port for this target.
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// (String) The target description.
	// The target description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Boolean expression to filter the workers used to access this target
	// Boolean expression to filter the workers used to access this target
	EgressWorkerFilter *string `json:"egressWorkerFilter,omitempty" tf:"egress_worker_filter,omitempty"`

	// (Boolean) HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	EnableSessionRecording *bool `json:"enableSessionRecording,omitempty" tf:"enable_session_recording,omitempty"`

	// (Set of String) A list of host source ID's. Cannot be used alongside address.
	// A list of host source ID's. Cannot be used alongside address.
	HostSourceIds []*string `json:"hostSourceIds,omitempty" tf:"host_source_ids,omitempty"`

	// (String) The ID of the target.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	// HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	IngressWorkerFilter *string `json:"ingressWorkerFilter,omitempty" tf:"ingress_worker_filter,omitempty"`

	// (Set of String) A list of injected application credential source ID's.
	// A list of injected application credential source ID's.
	InjectedApplicationCredentialSourceIds []*string `json:"injectedApplicationCredentialSourceIds,omitempty" tf:"injected_application_credential_source_ids,omitempty"`

	// (String) The target name. Defaults to the resource name.
	// The target name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (Number)
	SessionConnectionLimit *float64 `json:"sessionConnectionLimit,omitempty" tf:"session_connection_limit,omitempty"`

	// (Number)
	SessionMaxSeconds *float64 `json:"sessionMaxSeconds,omitempty" tf:"session_max_seconds,omitempty"`

	// (String) HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	StorageBucketID *string `json:"storageBucketId,omitempty" tf:"storage_bucket_id,omitempty"`

	// (String) The target resource type.
	// The target resource type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String, Deprecated) Boolean expression to filter the workers for this target
	// Boolean expression to filter the workers for this target
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

type TargetParameters struct {

	// (String) Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	// Optionally, a valid network address to connect to for this target. Cannot be used alongside host_source_ids.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (Set of String) A list of brokered credential source ID's.
	// A list of brokered credential source ID's.
	// +kubebuilder:validation:Optional
	BrokeredCredentialSourceIds []*string `json:"brokeredCredentialSourceIds,omitempty" tf:"brokered_credential_source_ids,omitempty"`

	// (Number) The default client port for this target.
	// The default client port for this target.
	// +kubebuilder:validation:Optional
	DefaultClientPort *float64 `json:"defaultClientPort,omitempty" tf:"default_client_port,omitempty"`

	// (Number) The default port for this target.
	// The default port for this target.
	// +kubebuilder:validation:Optional
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// (String) The target description.
	// The target description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Boolean expression to filter the workers used to access this target
	// Boolean expression to filter the workers used to access this target
	// +kubebuilder:validation:Optional
	EgressWorkerFilter *string `json:"egressWorkerFilter,omitempty" tf:"egress_worker_filter,omitempty"`

	// (Boolean) HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Enable sessions recording for this target. Only applicable for SSH targets.
	// +kubebuilder:validation:Optional
	EnableSessionRecording *bool `json:"enableSessionRecording,omitempty" tf:"enable_session_recording,omitempty"`

	// (Set of String) A list of host source ID's. Cannot be used alongside address.
	// A list of host source ID's. Cannot be used alongside address.
	// +kubebuilder:validation:Optional
	HostSourceIds []*string `json:"hostSourceIds,omitempty" tf:"host_source_ids,omitempty"`

	// (String) HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	// HCP Only. Boolean expression to filter the workers a user will connect to when initiating a session against this target
	// +kubebuilder:validation:Optional
	IngressWorkerFilter *string `json:"ingressWorkerFilter,omitempty" tf:"ingress_worker_filter,omitempty"`

	// (Set of String) A list of injected application credential source ID's.
	// A list of injected application credential source ID's.
	// +kubebuilder:validation:Optional
	InjectedApplicationCredentialSourceIds []*string `json:"injectedApplicationCredentialSourceIds,omitempty" tf:"injected_application_credential_source_ids,omitempty"`

	// (String) The target name. Defaults to the resource name.
	// The target name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID in which the resource is created. Defaults to the provider's default_scope if unset.
	// The scope ID in which the resource is created. Defaults to the provider's `default_scope` if unset.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	SessionConnectionLimit *float64 `json:"sessionConnectionLimit,omitempty" tf:"session_connection_limit,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	SessionMaxSeconds *float64 `json:"sessionMaxSeconds,omitempty" tf:"session_max_seconds,omitempty"`

	// (String) HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	// HCP/Ent Only. Storage bucket for this target. Only applicable for SSH targets.
	// +kubebuilder:validation:Optional
	StorageBucketID *string `json:"storageBucketId,omitempty" tf:"storage_bucket_id,omitempty"`

	// (String) The target resource type.
	// The target resource type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String, Deprecated) Boolean expression to filter the workers for this target
	// Boolean expression to filter the workers for this target
	// +kubebuilder:validation:Optional
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

// TargetSpec defines the desired state of Target
type TargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetParameters `json:"forProvider"`
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
	InitProvider TargetInitParameters `json:"initProvider,omitempty"`
}

// TargetStatus defines the observed state of Target.
type TargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Target is the Schema for the Targets API. The target resource allows you to configure a Boundary target.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   TargetSpec   `json:"spec"`
	Status TargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetList contains a list of Targets
type TargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Target `json:"items"`
}

// Repository type metadata.
var (
	Target_Kind             = "Target"
	Target_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Target_Kind}.String()
	Target_KindAPIVersion   = Target_Kind + "." + CRDGroupVersion.String()
	Target_GroupVersionKind = CRDGroupVersion.WithKind(Target_Kind)
)

func init() {
	SchemeBuilder.Register(&Target{}, &TargetList{})
}
