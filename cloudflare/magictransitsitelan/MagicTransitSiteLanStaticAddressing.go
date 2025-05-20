// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_lan#address MagicTransitSiteLan#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_lan#dhcp_relay MagicTransitSiteLan#dhcp_relay}.
	DhcpRelay *MagicTransitSiteLanStaticAddressingDhcpRelay `field:"optional" json:"dhcpRelay" yaml:"dhcpRelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_lan#dhcp_server MagicTransitSiteLan#dhcp_server}.
	DhcpServer *MagicTransitSiteLanStaticAddressingDhcpServer `field:"optional" json:"dhcpServer" yaml:"dhcpServer"`
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_lan#secondary_address MagicTransitSiteLan#secondary_address}
	SecondaryAddress *string `field:"optional" json:"secondaryAddress" yaml:"secondaryAddress"`
	// A valid CIDR notation representing an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_lan#virtual_address MagicTransitSiteLan#virtual_address}
	VirtualAddress *string `field:"optional" json:"virtualAddress" yaml:"virtualAddress"`
}

