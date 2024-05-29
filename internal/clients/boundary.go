/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/releaseband/crossplane-provider-boundary/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal boundary credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{
			// required
			"addr": creds["addr"], // (String) The base url of the Boundary API, e.g. "http://127.0.0.1:9200". If not set, it will be read from the "BOUNDARY_ADDR" env var.
		}

		// optional
		addIfSet(ps.Configuration, creds, "auth_method_id")                  // (String) The auth method ID e.g. ampw_1234567890. If not set, the default auth method for the given scope ID will be used.
		addIfSet(ps.Configuration, creds, "auth_method_login_name")          // (String) The auth method login name for password-style or ldap-style auth methods
		addIfSet(ps.Configuration, creds, "auth_method_password")            // (String) The auth method password for password-style or ldap-style auth methods
		addIfSet(ps.Configuration, creds, "password_auth_method_login_name") // (String, Deprecated) The auth method login name for password-style auth methods
		addIfSet(ps.Configuration, creds, "password_auth_method_password")   // (String, Deprecated) The auth method password for password-style auth methods
		addIfSet(ps.Configuration, creds, "plugin_execution_dir")            // (String) Specifies a directory that the Boundary provider can use to write and execute its built-in plugins.
		addIfSet(ps.Configuration, creds, "recovery_kms_hcl")                //  (String) Can be a heredoc string or a path on disk. If set, the string/file will be parsed as HCL and used with the recovery KMS mechanism. While this is set, it will override any other authentication information; the KMS mechanism will always be used. See Boundary's KMS docs for examples: https://boundaryproject.io/docs/configuration/kms
		addIfSet(ps.Configuration, creds, "scope_id")                        //  (String) The scope ID for the default auth method.
		addIfSet(ps.Configuration, creds, "tls_insecure")                    //  (Boolean) When set to true, does not validate the Boundary API endpoint certificate
		addIfSet(ps.Configuration, creds, "token")                           //  (String) The Boundary token to use, as a string or path on disk containing just the string. If set, the token read here will be used in place of authenticating with the auth method specified in "auth_method_id", although the recovery KMS mechanism will still override this. Can also be set with the BOUNDARY_TOKEN environment variable.

		return ps, nil
	}
}

func addIfSet(m map[string]any, creds map[string]string, key string) {
	if value, ok := creds[key]; ok {
		m[key] = value
	}
}
