package gmail

import (
	"context"
	"fmt"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("googleworkspace_gmail_alias", func(r *config.Resource) {
		r.ShortGroup = "gmail"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
			pe, ok := parameters["primary_email"]
			if !ok {
				return "", fmt.Errorf("primary_email not found in parameters")
			}
			sae, ok := parameters["send_as_email"]
			if !ok {
				return "", fmt.Errorf("send_as_email not found in parameters")
			}
			return fmt.Sprintf("%v:%v", pe, sae), nil
		}
	})
}
