/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"boundary_account":                                  config.IdentifierFromProvider,
	"boundary_account_ldap":                             config.IdentifierFromProvider,
	"boundary_account_oidc":                             config.IdentifierFromProvider,
	"boundary_account_password":                         config.IdentifierFromProvider,
	"boundary_alias_target":                             config.IdentifierFromProvider,
	"boundary_auth_method":                              config.IdentifierFromProvider,
	"boundary_auth_method_ldap":                         config.IdentifierFromProvider,
	"boundary_auth_method_oidc":                         config.IdentifierFromProvider,
	"boundary_auth_method_password":                     config.IdentifierFromProvider,
	"boundary_credential_json":                          config.IdentifierFromProvider,
	"boundary_credential_library_vault":                 config.IdentifierFromProvider,
	"boundary_credential_library_vault_ssh_certificate": config.IdentifierFromProvider,
	"boundary_credential_ssh_private_key":               config.IdentifierFromProvider,
	"boundary_credential_store_static":                  config.IdentifierFromProvider,
	"boundary_credential_store_vault":                   config.IdentifierFromProvider,
	"boundary_credential_username_password":             config.IdentifierFromProvider,
	"boundary_group":                                    config.IdentifierFromProvider,
	"boundary_host":                                     config.IdentifierFromProvider,
	"boundary_host_catalog":                             config.IdentifierFromProvider,
	"boundary_host_catalog_plugin":                      config.IdentifierFromProvider,
	"boundary_host_catalog_static":                      config.IdentifierFromProvider,
	"boundary_host_set":                                 config.IdentifierFromProvider,
	"boundary_host_set_plugin":                          config.IdentifierFromProvider,
	"boundary_host_set_static":                          config.IdentifierFromProvider,
	"boundary_host_static":                              config.IdentifierFromProvider,
	"boundary_managed_group":                            config.IdentifierFromProvider,
	"boundary_managed_group_ldap":                       config.IdentifierFromProvider,
	"boundary_policy_storage":                           config.IdentifierFromProvider,
	"boundary_role":                                     config.IdentifierFromProvider,
	"boundary_scope":                                    config.IdentifierFromProvider,
	"boundary_scope_policy_attachment":                  config.IdentifierFromProvider,
	"boundary_storage_bucket":                           config.IdentifierFromProvider,
	"boundary_target":                                   config.IdentifierFromProvider,
	"boundary_user":                                     config.IdentifierFromProvider,
	"boundary_worker":                                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
