// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationSaasAppCustomClaimsSource struct {
	// The name of the IdP claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A mapping from IdP ID to claim name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_application#name_by_idp ZeroTrustAccessApplication#name_by_idp}
	NameByIdp *map[string]*string `field:"optional" json:"nameByIdp" yaml:"nameByIdp"`
}

