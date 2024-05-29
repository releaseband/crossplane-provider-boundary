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

type OidcInitParameters struct {

	// The resource ID for the auth method.
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// The account description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The OIDC issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The account name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OIDC subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type OidcObservation struct {

	// The resource ID for the auth method.
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// The account description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The OIDC issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The account name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OIDC subject.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type OidcParameters struct {

	// The resource ID for the auth method.
	// +kubebuilder:validation:Optional
	AuthMethodID *string `json:"authMethodId,omitempty" tf:"auth_method_id,omitempty"`

	// The account description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The OIDC issuer.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The account name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The OIDC subject.
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`
}

// OidcSpec defines the desired state of Oidc
type OidcSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OidcParameters `json:"forProvider"`
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
	InitProvider OidcInitParameters `json:"initProvider,omitempty"`
}

// OidcStatus defines the observed state of Oidc.
type OidcStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OidcObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Oidc is the Schema for the Oidcs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Oidc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authMethodId) || (has(self.initProvider) && has(self.initProvider.authMethodId))",message="spec.forProvider.authMethodId is a required parameter"
	Spec   OidcSpec   `json:"spec"`
	Status OidcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OidcList contains a list of Oidcs
type OidcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Oidc `json:"items"`
}

// Repository type metadata.
var (
	Oidc_Kind             = "Oidc"
	Oidc_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Oidc_Kind}.String()
	Oidc_KindAPIVersion   = Oidc_Kind + "." + CRDGroupVersion.String()
	Oidc_GroupVersionKind = CRDGroupVersion.WithKind(Oidc_Kind)
)

func init() {
	SchemeBuilder.Register(&Oidc{}, &OidcList{})
}