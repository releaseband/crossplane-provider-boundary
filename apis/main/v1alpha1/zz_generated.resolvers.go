// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Scope.
func (mg *Scope) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ScopeID),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.ForProvider.ScopeIDRef,
		Selector:     mg.Spec.ForProvider.ScopeIDSelector,
		To: reference.To{
			List:    &ScopeList{},
			Managed: &Scope{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ScopeID")
	}
	mg.Spec.ForProvider.ScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ScopeID),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.InitProvider.ScopeIDRef,
		Selector:     mg.Spec.InitProvider.ScopeIDSelector,
		To: reference.To{
			List:    &ScopeList{},
			Managed: &Scope{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ScopeID")
	}
	mg.Spec.InitProvider.ScopeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ScopeIDRef = rsp.ResolvedReference

	return nil
}
