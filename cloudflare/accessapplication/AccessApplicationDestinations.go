// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationDestinations struct {
	// The URI of the destination.
	//
	// Public destinations can include a domain and path with wildcards. Private destinations are an early access feature and gated behind a feature flag. Private destinations support private IPv4, IPv6, and Server Name Indications (SNI) with optional port ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/access_application#uri AccessApplication#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The destination type. Available values: `public`, `private`. Defaults to `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/access_application#type AccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

