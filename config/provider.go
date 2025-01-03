/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	authconfig "github.com/releaseband/crossplane-provider-boundary/config/auth"
	boundary "github.com/releaseband/crossplane-provider-boundary/config/boundary"
	host "github.com/releaseband/crossplane-provider-boundary/config/host"
	mainconfig "github.com/releaseband/crossplane-provider-boundary/config/main"
	managed "github.com/releaseband/crossplane-provider-boundary/config/managed"
)

const (
	resourcePrefix = "boundary"
	modulePath     = "github.com/releaseband/crossplane-provider-boundary"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("boundary.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// repository.Configure,
		mainconfig.Configure,
		boundary.Configure,
		host.Configure,
		authconfig.Configure,
		managed.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
