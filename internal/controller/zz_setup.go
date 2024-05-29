// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ldap "github.com/releaseband/crossplane-provider-boundary/internal/controller/account/ldap"
	oidc "github.com/releaseband/crossplane-provider-boundary/internal/controller/account/oidc"
	password "github.com/releaseband/crossplane-provider-boundary/internal/controller/account/password"
	target "github.com/releaseband/crossplane-provider-boundary/internal/controller/alias/target"
	method "github.com/releaseband/crossplane-provider-boundary/internal/controller/auth/method"
	methodldap "github.com/releaseband/crossplane-provider-boundary/internal/controller/auth/methodldap"
	methodoidc "github.com/releaseband/crossplane-provider-boundary/internal/controller/auth/methodoidc"
	methodpassword "github.com/releaseband/crossplane-provider-boundary/internal/controller/auth/methodpassword"
	account "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/account"
	group "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/group"
	host "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/host"
	role "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/role"
	scope "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/scope"
	targetboundary "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/target"
	user "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/user"
	worker "github.com/releaseband/crossplane-provider-boundary/internal/controller/boundary/worker"
	json "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/json"
	libraryvault "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/libraryvault"
	libraryvaultsshcertificate "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/libraryvaultsshcertificate"
	sshprivatekey "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/sshprivatekey"
	storestatic "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/storestatic"
	storevault "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/storevault"
	usernamepassword "github.com/releaseband/crossplane-provider-boundary/internal/controller/credential/usernamepassword"
	catalog "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/catalog"
	catalogplugin "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/catalogplugin"
	catalogstatic "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/catalogstatic"
	set "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/set"
	setplugin "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/setplugin"
	setstatic "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/setstatic"
	static "github.com/releaseband/crossplane-provider-boundary/internal/controller/host/static"
	groupmanaged "github.com/releaseband/crossplane-provider-boundary/internal/controller/managed/group"
	groupldap "github.com/releaseband/crossplane-provider-boundary/internal/controller/managed/groupldap"
	storage "github.com/releaseband/crossplane-provider-boundary/internal/controller/policy/storage"
	providerconfig "github.com/releaseband/crossplane-provider-boundary/internal/controller/providerconfig"
	policyattachment "github.com/releaseband/crossplane-provider-boundary/internal/controller/scope/policyattachment"
	bucket "github.com/releaseband/crossplane-provider-boundary/internal/controller/storage/bucket"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ldap.Setup,
		oidc.Setup,
		password.Setup,
		target.Setup,
		method.Setup,
		methodldap.Setup,
		methodoidc.Setup,
		methodpassword.Setup,
		account.Setup,
		group.Setup,
		host.Setup,
		role.Setup,
		scope.Setup,
		targetboundary.Setup,
		user.Setup,
		worker.Setup,
		json.Setup,
		libraryvault.Setup,
		libraryvaultsshcertificate.Setup,
		sshprivatekey.Setup,
		storestatic.Setup,
		storevault.Setup,
		usernamepassword.Setup,
		catalog.Setup,
		catalogplugin.Setup,
		catalogstatic.Setup,
		set.Setup,
		setplugin.Setup,
		setstatic.Setup,
		static.Setup,
		groupmanaged.Setup,
		groupldap.Setup,
		storage.Setup,
		providerconfig.Setup,
		policyattachment.Setup,
		bucket.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
