// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessidentityprovider


type AccessIdentityProviderScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#enabled AccessIdentityProvider#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Deprecated. Use `identity_update_behavior`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#group_member_deprovision AccessIdentityProvider#group_member_deprovision}
	GroupMemberDeprovision interface{} `field:"optional" json:"groupMemberDeprovision" yaml:"groupMemberDeprovision"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	//
	// Use "automatic" to automatically update a user's identity and augment it with fields from the SCIM user resource. Use "reauth" to force re-authentication on group membership updates, user identity update will only occur after successful re-authentication. With "reauth" identities will not contain fields from the SCIM user resource. With "no_action" identities will not be changed by SCIM updates in any way and users will not be prompted to reauthenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#identity_update_behavior AccessIdentityProvider#identity_update_behavior}
	IdentityUpdateBehavior *string `field:"optional" json:"identityUpdateBehavior" yaml:"identityUpdateBehavior"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.
	//
	// This cannot be enabled unless user_deprovision is also enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#seat_deprovision AccessIdentityProvider#seat_deprovision}
	SeatDeprovision interface{} `field:"optional" json:"seatDeprovision" yaml:"seatDeprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first time.
	//
	// It is redacted on subsequent requests.  If you lose this you will need to refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#secret AccessIdentityProvider#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_identity_provider#user_deprovision AccessIdentityProvider#user_deprovision}
	UserDeprovision interface{} `field:"optional" json:"userDeprovision" yaml:"userDeprovision"`
}

