package chrome

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("googleworkspace_chrome_policy", func(r *config.Resource) {
		r.Kind = "ChromePolicy"
		r.ShortGroup = "chrome"
		r.References["org_unit_id"] = config.Reference{
			Type:              "github.com/alanraison/provider-googleworkspace/apis/directory/v1alpha1.OrgUnit",
			RefFieldName:      "OrgUnitRef",
			SelectorFieldName: "OrgUnitSelector",
		}
	})
}
