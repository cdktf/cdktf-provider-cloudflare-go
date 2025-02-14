// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan


type MagicTransitSiteLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/magic_transit_site_lan#server_addresses MagicTransitSiteLan#server_addresses}
	ServerAddresses *[]*string `field:"optional" json:"serverAddresses" yaml:"serverAddresses"`
}

