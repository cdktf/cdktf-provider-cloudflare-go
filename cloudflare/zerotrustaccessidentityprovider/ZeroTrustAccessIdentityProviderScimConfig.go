// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider


type ZeroTrustAccessIdentityProviderScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_identity_provider#enabled ZeroTrustAccessIdentityProvider#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates how a SCIM event updates a user identity used for policy evaluation.
	//
	// Use "automatic" to automatically update a user's identity and augment it with fields from the SCIM user resource. Use "reauth" to force re-authentication on group membership updates, user identity update will only occur after successful re-authentication. With "reauth" identities will not contain fields from the SCIM user resource. With "no_action" identities will not be changed by SCIM updates in any way and users will not be prompted to reauthenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_identity_provider#identity_update_behavior ZeroTrustAccessIdentityProvider#identity_update_behavior}
	IdentityUpdateBehavior *string `field:"optional" json:"identityUpdateBehavior" yaml:"identityUpdateBehavior"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned in the Identity Provider.
	//
	// This cannot be enabled unless user_deprovision is also enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_identity_provider#seat_deprovision ZeroTrustAccessIdentityProvider#seat_deprovision}
	SeatDeprovision interface{} `field:"optional" json:"seatDeprovision" yaml:"seatDeprovision"`
	// A flag to enable revoking a user's session in Access and Gateway when they have been deprovisioned in the Identity Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_identity_provider#user_deprovision ZeroTrustAccessIdentityProvider#user_deprovision}
	UserDeprovision interface{} `field:"optional" json:"userDeprovision" yaml:"userDeprovision"`
}

