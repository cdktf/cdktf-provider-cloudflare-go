// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasAppCustomClaimSource struct {
	// The name of the attribute as provided by the IDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A mapping from IdP ID to claim name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_application#name_by_idp AccessApplication#name_by_idp}
	NameByIdp *map[string]*string `field:"optional" json:"nameByIdp" yaml:"nameByIdp"`
}

