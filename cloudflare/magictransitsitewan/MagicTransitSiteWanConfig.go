// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsitewan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteWanConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#account_id MagicTransitSiteWan#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#physport MagicTransitSiteWan#physport}.
	Physport *float64 `field:"required" json:"physport" yaml:"physport"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#site_id MagicTransitSiteWan#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// VLAN port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#vlan_tag MagicTransitSiteWan#vlan_tag}
	VlanTag *float64 `field:"required" json:"vlanTag" yaml:"vlanTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#name MagicTransitSiteWan#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#priority MagicTransitSiteWan#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#static_addressing MagicTransitSiteWan#static_addressing}
	StaticAddressing *MagicTransitSiteWanStaticAddressing `field:"optional" json:"staticAddressing" yaml:"staticAddressing"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/magic_transit_site_wan#wan_id MagicTransitSiteWan#wan_id}
	WanId *string `field:"optional" json:"wanId" yaml:"wanId"`
}

