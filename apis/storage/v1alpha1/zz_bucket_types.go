// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type BucketInitParameters struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The name of the bucket within the external object store service.
	// The name of the bucket within the external object store service.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) The prefix used to organize the data held within the external object store.
	// The prefix used to organize the data held within the external object store.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// (String) The storage bucket description.
	// The storage bucket description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The storage bucket name. Defaults to the resource name.
	// The storage bucket name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// (String) The name of the plugin that should back the resource. This or plugin_id must be defined.
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	PluginName *string `json:"pluginName,omitempty" tf:"plugin_name,omitempty"`

	// (String) The scope for this storage bucket.
	// The scope for this storage bucket.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (String) Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

type BucketObservation struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The name of the bucket within the external object store service.
	// The name of the bucket within the external object store service.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) The prefix used to organize the data held within the external object store.
	// The prefix used to organize the data held within the external object store.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// (String) The storage bucket description.
	// The storage bucket description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the storage bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Internal only. Used to force update so that we can always check the value of secrets.
	// Internal only. Used to force update so that we can always check the value of secrets.
	InternalForceUpdate *string `json:"internalForceUpdate,omitempty" tf:"internal_force_update,omitempty"`

	// provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	InternalHMACUsedForSecretsConfigHMAC *string `json:"internalHmacUsedForSecretsConfigHmac,omitempty" tf:"internal_hmac_used_for_secrets_config_hmac,omitempty"`

	// (String) Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	InternalSecretsConfigHMAC *string `json:"internalSecretsConfigHmac,omitempty" tf:"internal_secrets_config_hmac,omitempty"`

	// (String) The storage bucket name. Defaults to the resource name.
	// The storage bucket name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// (String) The name of the plugin that should back the resource. This or plugin_id must be defined.
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	PluginName *string `json:"pluginName,omitempty" tf:"plugin_name,omitempty"`

	// (String) The scope for this storage bucket.
	// The scope for this storage bucket.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (String) The HMAC'd secrets value returned from the server.
	// The HMAC'd secrets value returned from the server.
	SecretsHMAC *string `json:"secretsHmac,omitempty" tf:"secrets_hmac,omitempty"`

	// (String) Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

type BucketParameters struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	// +kubebuilder:validation:Optional
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The name of the bucket within the external object store service.
	// The name of the bucket within the external object store service.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// (String) The prefix used to organize the data held within the external object store.
	// The prefix used to organize the data held within the external object store.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// (String) The storage bucket description.
	// The storage bucket description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The storage bucket name. Defaults to the resource name.
	// The storage bucket name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// +kubebuilder:validation:Optional
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// (String) The name of the plugin that should back the resource. This or plugin_id must be defined.
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	// +kubebuilder:validation:Optional
	PluginName *string `json:"pluginName,omitempty" tf:"plugin_name,omitempty"`

	// (String) The scope for this storage bucket.
	// The scope for this storage bucket.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	// +kubebuilder:validation:Optional
	SecretsJSONSecretRef *v1.SecretKeySelector `json:"secretsJsonSecretRef,omitempty" tf:"-"`

	// (String) Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	// +kubebuilder:validation:Optional
	WorkerFilter *string `json:"workerFilter,omitempty" tf:"worker_filter,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. The storage bucket resource allows you to configure a Boundary storage bucket. A storage bucket can only belong to the Global scope or an Org scope. At this time, the only supported storage for storage buckets is AWS S3. This feature requires Boundary Enterprise or Boundary HCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucketName) || (has(self.initProvider) && has(self.initProvider.bucketName))",message="spec.forProvider.bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.workerFilter) || (has(self.initProvider) && has(self.initProvider.workerFilter))",message="spec.forProvider.workerFilter is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
