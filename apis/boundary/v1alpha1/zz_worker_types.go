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

type WorkerInitParameters struct {

	// (String) The description for the worker.
	// The description for the worker.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name for the worker.
	// The name for the worker.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope for the worker. Defaults to global.
	// The scope for the worker. Defaults to `global`.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// led authentication flow. Leaving this blank will result in a controller generated token.
	// The worker authentication token required to register the worker for the worker-led authentication flow. Leaving this blank will result in a controller generated token.
	WorkerGeneratedAuthToken *string `json:"workerGeneratedAuthToken,omitempty" tf:"worker_generated_auth_token,omitempty"`
}

type WorkerObservation struct {

	// (String) The accessible address of the self managed worker.
	// The accessible address of the self managed worker.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (List of String) A list of actions that the worker is entitled to perform.
	// A list of actions that the worker is entitled to perform.
	AuthorizedActions []*string `json:"authorizedActions,omitempty" tf:"authorized_actions,omitempty"`

	// managed worker.
	// A single use token generated by the controller to be passed to the self-managed worker.
	ControllerGeneratedActivationToken *string `json:"controllerGeneratedActivationToken,omitempty" tf:"controller_generated_activation_token,omitempty"`

	// (String) The description for the worker.
	// The description for the worker.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of the worker.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name for the worker.
	// The name for the worker.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The version of the Boundary binary running on the self managed worker.
	// The version of the Boundary binary running on the self managed worker.
	ReleaseVersion *float64 `json:"releaseVersion,omitempty" tf:"release_version,omitempty"`

	// (String) The scope for the worker. Defaults to global.
	// The scope for the worker. Defaults to `global`.
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// led authentication flow. Leaving this blank will result in a controller generated token.
	// The worker authentication token required to register the worker for the worker-led authentication flow. Leaving this blank will result in a controller generated token.
	WorkerGeneratedAuthToken *string `json:"workerGeneratedAuthToken,omitempty" tf:"worker_generated_auth_token,omitempty"`
}

type WorkerParameters struct {

	// (String) The description for the worker.
	// The description for the worker.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name for the worker.
	// The name for the worker.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The scope for the worker. Defaults to global.
	// The scope for the worker. Defaults to `global`.
	// +kubebuilder:validation:Optional
	ScopeID *string `json:"scopeId,omitempty" tf:"scope_id,omitempty"`

	// led authentication flow. Leaving this blank will result in a controller generated token.
	// The worker authentication token required to register the worker for the worker-led authentication flow. Leaving this blank will result in a controller generated token.
	// +kubebuilder:validation:Optional
	WorkerGeneratedAuthToken *string `json:"workerGeneratedAuthToken,omitempty" tf:"worker_generated_auth_token,omitempty"`
}

// WorkerSpec defines the desired state of Worker
type WorkerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkerParameters `json:"forProvider"`
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
	InitProvider WorkerInitParameters `json:"initProvider,omitempty"`
}

// WorkerStatus defines the observed state of Worker.
type WorkerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Worker is the Schema for the Workers API. The resource allows you to create a self-managed worker object.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,boundary}
type Worker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkerSpec   `json:"spec"`
	Status            WorkerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkerList contains a list of Workers
type WorkerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Worker `json:"items"`
}

// Repository type metadata.
var (
	Worker_Kind             = "Worker"
	Worker_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Worker_Kind}.String()
	Worker_KindAPIVersion   = Worker_Kind + "." + CRDGroupVersion.String()
	Worker_GroupVersionKind = CRDGroupVersion.WithKind(Worker_Kind)
)

func init() {
	SchemeBuilder.Register(&Worker{}, &WorkerList{})
}
