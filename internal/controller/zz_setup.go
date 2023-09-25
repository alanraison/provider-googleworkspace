/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	alias "github.com/alanraison/provider-googleworkspace/internal/controller/directory/alias"
	assignment "github.com/alanraison/provider-googleworkspace/internal/controller/directory/assignment"
	domain "github.com/alanraison/provider-googleworkspace/internal/controller/directory/domain"
	group "github.com/alanraison/provider-googleworkspace/internal/controller/directory/group"
	member "github.com/alanraison/provider-googleworkspace/internal/controller/directory/member"
	members "github.com/alanraison/provider-googleworkspace/internal/controller/directory/members"
	orgunit "github.com/alanraison/provider-googleworkspace/internal/controller/directory/orgunit"
	role "github.com/alanraison/provider-googleworkspace/internal/controller/directory/role"
	schema "github.com/alanraison/provider-googleworkspace/internal/controller/directory/schema"
	user "github.com/alanraison/provider-googleworkspace/internal/controller/directory/user"
	settings "github.com/alanraison/provider-googleworkspace/internal/controller/groups/settings"
	providerconfig "github.com/alanraison/provider-googleworkspace/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alias.Setup,
		assignment.Setup,
		domain.Setup,
		group.Setup,
		member.Setup,
		members.Setup,
		orgunit.Setup,
		role.Setup,
		schema.Setup,
		user.Setup,
		settings.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
