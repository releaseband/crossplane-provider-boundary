/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/releaseband/crossplane-provider-boundary/apis/managed/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Role.
func (mg *Role) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.PrincipalIds),
		Extract:       resource.ExtractParamPath("id", true),
		References:    mg.Spec.ForProvider.PrincipalIdsRefs,
		Selector:      mg.Spec.ForProvider.PrincipalIdsSelector,
		To: reference.To{
			List:    &v1alpha1.GroupList{},
			Managed: &v1alpha1.Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalIds")
	}
	mg.Spec.ForProvider.PrincipalIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.PrincipalIdsRefs = mrsp.ResolvedReferences

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

	return nil
}

// ResolveReferences of this Target.
func (mg *Target) ResolveReferences(ctx context.Context, c client.Reader) error {
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

	return nil
}
