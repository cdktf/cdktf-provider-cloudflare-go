// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessidentityprovider


type ZeroTrustAccessIdentityProviderScimConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_identity_provider#enabled ZeroTrustAccessIdentityProvider#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_identity_provider#group_member_deprovision ZeroTrustAccessIdentityProvider#group_member_deprovision}.
	GroupMemberDeprovision interface{} `field:"optional" json:"groupMemberDeprovision" yaml:"groupMemberDeprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_identity_provider#seat_deprovision ZeroTrustAccessIdentityProvider#seat_deprovision}.
	SeatDeprovision interface{} `field:"optional" json:"seatDeprovision" yaml:"seatDeprovision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_identity_provider#secret ZeroTrustAccessIdentityProvider#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_identity_provider#user_deprovision ZeroTrustAccessIdentityProvider#user_deprovision}.
	UserDeprovision interface{} `field:"optional" json:"userDeprovision" yaml:"userDeprovision"`
}

