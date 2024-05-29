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

type JSONInitParameters struct {

	// (String) The credential store in which to save this json credential.
	// The credential store in which to save this json credential.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this json credential.
	// The description of this json credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of this json credential. Defaults to the resource name.
	// The name of this json credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type JSONObservation struct {

	// (String) The credential store in which to save this json credential.
	// The credential store in which to save this json credential.
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this json credential.
	// The description of this json credential.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this json credential.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of this json credential. Defaults to the resource name.
	// The name of this json credential. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The object hmac.
	// The object hmac.
	ObjectHMAC *string `json:"objectHmac,omitempty" tf:"object_hmac,omitempty"`
}

type JSONParameters struct {

	// (String) The credential store in which to save this json credential.
	// The credential store in which to save this json credential.
	// +kubebuilder:validation:Optional
	CredentialStoreID *string `json:"credentialStoreId,omitempty" tf:"credential_store_id,omitempty"`

	// (String) The description of this json credential.
	// The description of this json credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of this json credential. Defaults to the resource name.
	// The name of this json credential. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// escaped JSON string, or a file
	// The object for the this json credential. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file
	// +kubebuilder:validation:Optional
	ObjectSecretRef v1.SecretKeySelector `json:"objectSecretRef" tf:"-"`
}

// JSONSpec defines the desired state of JSON
type JSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JSONParameters `json:"forProvider"`
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
	InitProvider JSONInitParameters `json:"initProvider,omitempty"`
}

// JSONStatus defines the observed state of JSON.
type JSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JSON is the Schema for the JSONs API. The json credential resource allows you to congiure a credential using a json object.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type JSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialStoreId) || (has(self.initProvider) && has(self.initProvider.credentialStoreId))",message="spec.forProvider.credentialStoreId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectSecretRef)",message="spec.forProvider.objectSecretRef is a required parameter"
	Spec   JSONSpec   `json:"spec"`
	Status JSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JSONList contains a list of JSONs
type JSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JSON `json:"items"`
}

// Repository type metadata.
var (
	JSON_Kind             = "JSON"
	JSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JSON_Kind}.String()
	JSON_KindAPIVersion   = JSON_Kind + "." + CRDGroupVersion.String()
	JSON_GroupVersionKind = CRDGroupVersion.WithKind(JSON_Kind)
)

func init() {
	SchemeBuilder.Register(&JSON{}, &JSONList{})
}