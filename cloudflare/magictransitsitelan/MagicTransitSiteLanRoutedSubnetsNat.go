// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/magic_transit_site_lan#static_prefix MagicTransitSiteLan#static_prefix}
	StaticPrefix *string `field:"optional" json:"staticPrefix" yaml:"staticPrefix"`
}

