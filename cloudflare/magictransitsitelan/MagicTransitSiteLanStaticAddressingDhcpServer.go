// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_lan#dhcp_pool_end MagicTransitSiteLan#dhcp_pool_end}
	DhcpPoolEnd *string `field:"optional" json:"dhcpPoolEnd" yaml:"dhcpPoolEnd"`
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_lan#dhcp_pool_start MagicTransitSiteLan#dhcp_pool_start}
	DhcpPoolStart *string `field:"optional" json:"dhcpPoolStart" yaml:"dhcpPoolStart"`
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_lan#dns_server MagicTransitSiteLan#dns_server}
	DnsServer *string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_lan#dns_servers MagicTransitSiteLan#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Mapping of MAC addresses to IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_lan#reservations MagicTransitSiteLan#reservations}
	Reservations *map[string]*string `field:"optional" json:"reservations" yaml:"reservations"`
}

