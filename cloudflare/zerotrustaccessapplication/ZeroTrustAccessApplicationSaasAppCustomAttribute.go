// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomAttribute struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_access_application#source ZeroTrustAccessApplication#source}
	Source *ZeroTrustAccessApplicationSaasAppCustomAttributeSource `field:"required" json:"source" yaml:"source"`
	// A friendly name for the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_access_application#friendly_name ZeroTrustAccessApplication#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// The name of the attribute as provided to the SaaS app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A globally unique name for an identity or service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_access_application#name_format ZeroTrustAccessApplication#name_format}
	NameFormat *string `field:"optional" json:"nameFormat" yaml:"nameFormat"`
	// True if the attribute must be always present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_access_application#required ZeroTrustAccessApplication#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

