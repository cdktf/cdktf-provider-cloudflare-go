// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitelan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteLanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#account_id MagicTransitSiteLan#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#physport MagicTransitSiteLan#physport}.
	Physport *float64 `field:"required" json:"physport" yaml:"physport"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#site_id MagicTransitSiteLan#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// mark true to use this LAN for HA probing.
	//
	// only works for site with HA turned on. only one LAN can be set as the ha_link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#ha_link MagicTransitSiteLan#ha_link}
	HaLink interface{} `field:"optional" json:"haLink" yaml:"haLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#name MagicTransitSiteLan#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#nat MagicTransitSiteLan#nat}.
	Nat *MagicTransitSiteLanNat `field:"optional" json:"nat" yaml:"nat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#routed_subnets MagicTransitSiteLan#routed_subnets}.
	RoutedSubnets interface{} `field:"optional" json:"routedSubnets" yaml:"routedSubnets"`
	// If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP).
	//
	// However, if in high availability mode, static_address is required along with secondary and virtual address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#static_addressing MagicTransitSiteLan#static_addressing}
	StaticAddressing *MagicTransitSiteLanStaticAddressing `field:"optional" json:"staticAddressing" yaml:"staticAddressing"`
	// VLAN ID. Use zero for untagged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/magic_transit_site_lan#vlan_tag MagicTransitSiteLan#vlan_tag}
	VlanTag *float64 `field:"optional" json:"vlanTag" yaml:"vlanTag"`
}

