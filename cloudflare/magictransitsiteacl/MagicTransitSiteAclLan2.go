// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsiteacl


type MagicTransitSiteAclLan2 struct {
	// The identifier for the LAN you want to create an ACL policy with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_acl#lan_id MagicTransitSiteAcl#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
	// The name of the LAN based on the provided lan_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_acl#lan_name MagicTransitSiteAcl#lan_name}
	LanName *string `field:"optional" json:"lanName" yaml:"lanName"`
	// Array of port ranges on the provided LAN that will be included in the ACL.
	//
	// If no ports or port rangess are provided, communication on any port on this LAN is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_acl#port_ranges MagicTransitSiteAcl#port_ranges}
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// Array of ports on the provided LAN that will be included in the ACL.
	//
	// If no ports or port ranges are provided, communication on any port on this LAN is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_acl#ports MagicTransitSiteAcl#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL.
	//
	// If no subnets are provided, communication on any subnets on this LAN are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site_acl#subnets MagicTransitSiteAcl#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

