// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanRoutedSubnets struct {
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_transit_site_lan#next_hop MagicTransitSiteLan#next_hop}
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_transit_site_lan#prefix MagicTransitSiteLan#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_transit_site_lan#nat MagicTransitSiteLan#nat}.
	Nat *MagicTransitSiteLanRoutedSubnetsNat `field:"optional" json:"nat" yaml:"nat"`
}

