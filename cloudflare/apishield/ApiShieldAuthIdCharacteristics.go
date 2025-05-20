// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apishield


type ApiShieldAuthIdCharacteristics struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/api_shield#name ApiShield#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of characteristic. Available values: "header", "cookie", "jwt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/api_shield#type ApiShield#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

