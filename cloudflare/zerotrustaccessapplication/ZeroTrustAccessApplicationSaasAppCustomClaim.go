// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomClaim struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#source ZeroTrustAccessApplication#source}
	Source *ZeroTrustAccessApplicationSaasAppCustomClaimSource `field:"required" json:"source" yaml:"source"`
	// The name of the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// True if the attribute must be always present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#required ZeroTrustAccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The scope of the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_application#scope ZeroTrustAccessApplication#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

