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

type CatalogPluginInitParameters struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The host catalog description.
	// The host catalog description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Internal only. Used to force update so that we can always check the value of secrets.
	// Internal only. Used to force update so that we can always check the value of secrets.
	InternalForceUpdate *string `json:"internalForceUpdate,omitempty" tf:"internal_force_update,omitempty"`

	// provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	InternalHMACUsedForSecretsConfigHMAC *string `json:"internalHmacUsedForSecretsConfigHmac,omitempty" tf:"internal_hmac_used_for_secrets_config_hmac,omitempty"`

	// (String) Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	InternalSecretsConfigHMAC *string `json:"internalSecretsConfigHmac,omitempty" tf:"internal_secrets_config_hmac,omitempty"`

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// (String) The name of the plugin that should back the resource. This or plugin_id must be defined.
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	PluginName *string `json:"pluginName,omitempty" tf:"plugin_name,omitempty"`

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (String) The HMAC'd secrets value returned from the server.
	// The HMAC'd secrets value returned from the server.
	SecretsHMAC *string `json:"secretsHmac,omitempty" tf:"secrets_hmac,omitempty"`
}

type CatalogPluginObservation struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The host catalog description.
	// The host catalog description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the host catalog.
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

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the plugin that should back the resource. This or plugin_name must be defined.
	// The ID of the plugin that should back the resource. This or plugin_name must be defined.
	PluginID *string `json:"pluginId,omitempty" tf:"plugin_id,omitempty"`

	// (String) The name of the plugin that should back the resource. This or plugin_id must be defined.
	// The name of the plugin that should back the resource. This or plugin_id must be defined.
	PluginName *string `json:"pluginName,omitempty" tf:"plugin_name,omitempty"`

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (String) The HMAC'd secrets value returned from the server.
	// The HMAC'd secrets value returned from the server.
	SecretsHMAC *string `json:"secretsHmac,omitempty" tf:"secrets_hmac,omitempty"`
}

type CatalogPluginParameters struct {

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
	// +kubebuilder:validation:Optional
	AttributesJSON *string `json:"attributesJson,omitempty" tf:"attributes_json,omitempty"`

	// (String) The host catalog description.
	// The host catalog description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Internal only. Used to force update so that we can always check the value of secrets.
	// Internal only. Used to force update so that we can always check the value of secrets.
	// +kubebuilder:validation:Optional
	InternalForceUpdate *string `json:"internalForceUpdate,omitempty" tf:"internal_force_update,omitempty"`

	// provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	// +kubebuilder:validation:Optional
	InternalHMACUsedForSecretsConfigHMAC *string `json:"internalHmacUsedForSecretsConfigHmac,omitempty" tf:"internal_hmac_used_for_secrets_config_hmac,omitempty"`

	// (String) Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	// +kubebuilder:validation:Optional
	InternalSecretsConfigHMAC *string `json:"internalSecretsConfigHmac,omitempty" tf:"internal_secrets_config_hmac,omitempty"`

	// (String) The host catalog name. Defaults to the resource name.
	// The host catalog name. Defaults to the resource name.
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

	// (String) The scope ID in which the resource is created.
	// The scope ID in which the resource is created.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// (String) The HMAC'd secrets value returned from the server.
	// The HMAC'd secrets value returned from the server.
	// +kubebuilder:validation:Optional
	SecretsHMAC *string `json:"secretsHmac,omitempty" tf:"secrets_hmac,omitempty"`

	// escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
	// The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them for storage.
	// +kubebuilder:validation:Optional
	SecretsJSONSecretRef *v1.SecretKeySelector `json:"secretsJsonSecretRef,omitempty" tf:"-"`
}

// CatalogPluginSpec defines the desired state of CatalogPlugin
type CatalogPluginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogPluginParameters `json:"forProvider"`
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
	InitProvider CatalogPluginInitParameters `json:"initProvider,omitempty"`
}

// CatalogPluginStatus defines the observed state of CatalogPlugin.
type CatalogPluginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogPluginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogPlugin is the Schema for the CatalogPlugins API. The host catalog resource allows you to configure a Boundary plugin-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type CatalogPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopeId) || (has(self.initProvider) && has(self.initProvider.scopeId))",message="spec.forProvider.scopeId is a required parameter"
	Spec   CatalogPluginSpec   `json:"spec"`
	Status CatalogPluginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogPluginList contains a list of CatalogPlugins
type CatalogPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogPlugin `json:"items"`
}

// Repository type metadata.
var (
	CatalogPlugin_Kind             = "CatalogPlugin"
	CatalogPlugin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogPlugin_Kind}.String()
	CatalogPlugin_KindAPIVersion   = CatalogPlugin_Kind + "." + CRDGroupVersion.String()
	CatalogPlugin_GroupVersionKind = CRDGroupVersion.WithKind(CatalogPlugin_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogPlugin{}, &CatalogPluginList{})
}
