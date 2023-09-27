/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"golang.org/x/oauth2/google"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/alanraison/provider-googleworkspace/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal googleworkspace credentials as JSON"

	accessToken           = "access_token"
	credentials           = "credentials"
	customerID            = "customer_id"
	impersonatedUserEmail = "impersonated_user_email"
	oauthScopes           = "oauth_scopes"
	serviceAccount        = "service_account"
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

		ps.Configuration = map[string]any{}

		switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			_, err := google.DefaultTokenSource(ctx, "https://www.googleapis.com/auth/admin.directory.groups", "https://www.googleapis.com/auth/admin.directory.users")
			if err != nil {
				return ps, errors.Wrap(err, "getting default credentials source")
			}
		default:
			data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}
			creds := map[string]string{}
			if err := json.Unmarshal(data, &creds); err != nil {
				return ps, errors.Wrap(err, errUnmarshalCredentials)
			}
			extractCredentials(creds, ps.Configuration)
		}
		ps.Configuration[customerID] = pc.Spec.CustomerID

		return ps, nil
	}
}

func extractCredentials(creds map[string]string, config map[string]any) {
	// Set credentials in Terraform provider configuration.
	if v, ok := creds[accessToken]; ok {
		config[accessToken] = v
	}
	if v, ok := creds[credentials]; ok {
		config[credentials] = v
	}
	if v, ok := creds[impersonatedUserEmail]; ok {
		config[impersonatedUserEmail] = v
	}
	if v, ok := creds[oauthScopes]; ok {
		config[oauthScopes] = v
	}
	if v, ok := creds[serviceAccount]; ok {
		config[serviceAccount] = v
	}
}
