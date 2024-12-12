// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomAttributeSource struct {
	// The name of the attribute as provided by the IDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A mapping from IdP ID to claim name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_access_application#name_by_idp ZeroTrustAccessApplication#name_by_idp}
	NameByIdp *map[string]*string `field:"optional" json:"nameByIdp" yaml:"nameByIdp"`
}

