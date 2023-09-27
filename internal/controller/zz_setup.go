/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	chromepolicy "github.com/alanraison/provider-googleworkspace/internal/controller/chrome/chromepolicy"
	domain "github.com/alanraison/provider-googleworkspace/internal/controller/directory/domain"
	domainalias "github.com/alanraison/provider-googleworkspace/internal/controller/directory/domainalias"
	group "github.com/alanraison/provider-googleworkspace/internal/controller/directory/group"
	groupmember "github.com/alanraison/provider-googleworkspace/internal/controller/directory/groupmember"
	groupmembers "github.com/alanraison/provider-googleworkspace/internal/controller/directory/groupmembers"
	orgunit "github.com/alanraison/provider-googleworkspace/internal/controller/directory/orgunit"
	role "github.com/alanraison/provider-googleworkspace/internal/controller/directory/role"
	roleassignment "github.com/alanraison/provider-googleworkspace/internal/controller/directory/roleassignment"
	schema "github.com/alanraison/provider-googleworkspace/internal/controller/directory/schema"
	user "github.com/alanraison/provider-googleworkspace/internal/controller/directory/user"
	settings "github.com/alanraison/provider-googleworkspace/internal/controller/groups/settings"
	providerconfig "github.com/alanraison/provider-googleworkspace/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		chromepolicy.Setup,
		domain.Setup,
		domainalias.Setup,
		group.Setup,
		groupmember.Setup,
		groupmembers.Setup,
		orgunit.Setup,
		role.Setup,
		roleassignment.Setup,
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
