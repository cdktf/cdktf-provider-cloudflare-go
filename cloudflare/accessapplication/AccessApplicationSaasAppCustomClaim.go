// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasAppCustomClaim struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/access_application#source AccessApplication#source}
	Source *AccessApplicationSaasAppCustomClaimSource `field:"required" json:"source" yaml:"source"`
	// The name of the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// True if the attribute must be always present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/access_application#required AccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The scope of the claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/access_application#scope AccessApplication#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

