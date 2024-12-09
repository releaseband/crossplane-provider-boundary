// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package host

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("boundary_host_catalog_static", func(r *config.Resource) {
		// r.References["scope_id"] = config.Reference{
		// 	TerraformName: "boundary_scope",
		// 	Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,

		// }
	})
	p.AddResourceConfigurator("boundary_host_static", func(r *config.Resource) {
		r.References["host_catalog_id"] = config.Reference{
			TerraformName: "boundary_host_catalog_static",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
	p.AddResourceConfigurator("boundary_host_set_static", func(r *config.Resource) {
		r.References["host_catalog_id"] = config.Reference{
			TerraformName: "boundary_host_catalog_static",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
		r.References["host_ids"] = config.Reference{
			TerraformName: "boundary_host_static",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
