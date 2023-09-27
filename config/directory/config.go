package directory

import "github.com/upbound/upjet/pkg/config"

const group = "directory"

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
		r.ShortGroup = group
		r.ExternalName = config.ParameterAsIdentifier("domain_name")
	})
	p.AddResourceConfigurator("googleworkspace_domain_alias", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "DomainAlias"
	})
	p.AddResourceConfigurator("googleworkspace_group", func(r *config.Resource) {
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("googleworkspace_group_member", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "GroupMember"
		r.References["group_id"] = config.Reference{
			Type:              "Group",
			RefFieldName:      "GroupRef",
			SelectorFieldName: "GroupSelector",
		}
	})
	p.AddResourceConfigurator("googleworkspace_group_members", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "GroupMembers"
		r.References["group_id"] = config.Reference{
			Type:              "Group",
			RefFieldName:      "GroupRef",
			SelectorFieldName: "GroupSelector",
		}
	})
	p.AddResourceConfigurator("googleworkspace_org_unit", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "OrgUnit"
	})
	p.AddResourceConfigurator("googleworkspace_role", func(r *config.Resource) {
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("googleworkspace_role_assignment", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "RoleAssignment"
		r.References["assigned_to"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
		r.References["role_id"] = config.Reference{
			Type:              "Role",
			RefFieldName:      "RoleRef",
			SelectorFieldName: "RoleSelector",
		}
		r.References["org_unit_id"] = config.Reference{
			Type:              "OrgUnit",
			RefFieldName:      "OrgUnitRef",
			SelectorFieldName: "OrgUnitSelector",
		}
	})
	p.AddResourceConfigurator("googleworkspace_user", func(r *config.Resource) {
		r.ShortGroup = group
		r.References["custom_schemas.schema_name"] = config.Reference{
			Type:              "Schema",
			RefFieldName:      "SchemaRef",
			SelectorFieldName: "SchemaSelector",
		}
	})
	p.AddResourceConfigurator("googleworkspace_schema", func(r *config.Resource) {
		r.ShortGroup = group
	})
}
