package directory

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// googleworkspace_domain
	// googleworkspace_domain_alias
	// googleworkspace_group
	// googleworkspace_group_member
	// googleworkspace_group_members
	// googleworkspace_org_unit
	// googleworkspace_role
	// googleworkspace_role_assignment
	// googleworkspace_user
	// googleworkspace_schema
	p.AddResourceConfigurator("googleworkspace_domain", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.ExternalName = config.ParameterAsIdentifier("domain_name")
	})
	p.AddResourceConfigurator("googleworkspace_domain_alias", func(r *config.Resource) {
		r.ShortGroup = "directory"
	})
	p.AddResourceConfigurator("googleworkspace_group", func(r *config.Resource) {
		r.ShortGroup = "directory"
	})
	p.AddResourceConfigurator("googleworkspace_group_member", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("googleworkspace_group_members", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("googleworkspace_org_unit", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.Kind = "OrgUnit"
	})
	p.AddResourceConfigurator("googleworkspace_role", func(r *config.Resource) {
		r.ShortGroup = "directory"
	})
	p.AddResourceConfigurator("googleworkspace_role_assignment", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.References["assigned_to"] = config.Reference{
			Type: "User",
		}
		r.References["role_id"] = config.Reference{
			Type: "Role",
		}
		r.References["org_unit_id"] = config.Reference{
			Type: "OrgUnit",
		}
	})
	p.AddResourceConfigurator("googleworkspace_user", func(r *config.Resource) {
		r.ShortGroup = "directory"
		r.References["custom_schemas.schema_name"] = config.Reference{
			Type: "Schema",
		}
	})
	p.AddResourceConfigurator("googleworkspace_schema", func(r *config.Resource) {
		r.ShortGroup = "directory"
	})
}
