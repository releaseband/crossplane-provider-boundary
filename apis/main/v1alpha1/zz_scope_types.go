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

type ScopeInitParameters struct {

	// (Boolean) If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	AutoCreateAdminRole *bool `json:"autoCreateAdminRole,omitempty" tf:"auto_create_admin_role,omitempty"`

	// (Boolean) Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	AutoCreateDefaultRole *bool `json:"autoCreateDefaultRole,omitempty" tf:"auto_create_default_role,omitempty"`

	// (String) The scope description.
	// The scope description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	GlobalScope *bool `json:"globalScope,omitempty" tf:"global_scope,omitempty"`

	// (String) The scope name. Defaults to the resource name.
	// The scope name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID containing the sub scope resource.
	// The scope ID containing the sub scope resource.
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

type ScopeObservation struct {

	// (Boolean) If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	AutoCreateAdminRole *bool `json:"autoCreateAdminRole,omitempty" tf:"auto_create_admin_role,omitempty"`

	// (Boolean) Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	AutoCreateDefaultRole *bool `json:"autoCreateDefaultRole,omitempty" tf:"auto_create_default_role,omitempty"`

	// (String) The scope description.
	// The scope description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	GlobalScope *bool `json:"globalScope,omitempty" tf:"global_scope,omitempty"`

	// (String) The ID of the scope.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The scope name. Defaults to the resource name.
	// The scope name. Defaults to the resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID containing the sub scope resource.
	// The scope ID containing the sub scope resource.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`
}

type ScopeParameters struct {

	// (Boolean) If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	// If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives permissions to manage the scope to the provider's user.
	// +kubebuilder:validation:Optional
	AutoCreateAdminRole *bool `json:"autoCreateAdminRole,omitempty" tf:"auto_create_admin_role,omitempty"`

	// (Boolean) Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	// Only relevant when creating an org scope. If set, when a new scope is created, the provider will not disable the functionality that automatically creates a role in the new scope and gives listing of scopes and auth methods and the ability to authenticate to the anonymous user.
	// +kubebuilder:validation:Optional
	AutoCreateDefaultRole *bool `json:"autoCreateDefaultRole,omitempty" tf:"auto_create_default_role,omitempty"`

	// (String) The scope description.
	// The scope description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	// Indicates that the scope containing this value is the global scope, which triggers some specialized behavior to allow it to be imported and managed.
	// +kubebuilder:validation:Optional
	GlobalScope *bool `json:"globalScope,omitempty" tf:"global_scope,omitempty"`

	// (String) The scope name. Defaults to the resource name.
	// The scope name. Defaults to the resource name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope ID containing the sub scope resource.
	// The scope ID containing the sub scope resource.
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

// ScopeSpec defines the desired state of Scope
type ScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScopeParameters `json:"forProvider"`
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
	InitProvider ScopeInitParameters `json:"initProvider,omitempty"`
}

// ScopeStatus defines the observed state of Scope.
type ScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Scope is the Schema for the Scopes API. The scope resource allows you to configure a Boundary scope.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScopeSpec   `json:"spec"`
	Status            ScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScopeList contains a list of Scopes
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scope `json:"items"`
}

// Repository type metadata.
var (
	Scope_Kind             = "Scope"
	Scope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Scope_Kind}.String()
	Scope_KindAPIVersion   = Scope_Kind + "." + CRDGroupVersion.String()
	Scope_GroupVersionKind = CRDGroupVersion.WithKind(Scope_Kind)
)

func init() {
	SchemeBuilder.Register(&Scope{}, &ScopeList{})
}
