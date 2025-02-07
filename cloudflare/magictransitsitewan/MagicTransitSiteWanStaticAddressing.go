// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitewan


type MagicTransitSiteWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#address MagicTransitSiteWan#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#gateway_address MagicTransitSiteWan#gateway_address}
	GatewayAddress *string `field:"required" json:"gatewayAddress" yaml:"gatewayAddress"`
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#secondary_address MagicTransitSiteWan#secondary_address}
	SecondaryAddress *string `field:"optional" json:"secondaryAddress" yaml:"secondaryAddress"`
}

