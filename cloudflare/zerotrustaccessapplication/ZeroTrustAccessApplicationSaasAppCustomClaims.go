// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomClaims struct {
	// The name of the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If the claim is required when building an OIDC token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#required ZeroTrustAccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The scope of the claim. Available values: "groups", "profile", "email", "openid".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#scope ZeroTrustAccessApplication#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#source ZeroTrustAccessApplication#source}.
	Source *ZeroTrustAccessApplicationSaasAppCustomClaimsSource `field:"optional" json:"source" yaml:"source"`
}

