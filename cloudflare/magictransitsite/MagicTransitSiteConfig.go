// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsite

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#account_id MagicTransitSite#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#name MagicTransitSite#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Magic Connector identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#connector_id MagicTransitSite#connector_id}
	ConnectorId *string `field:"optional" json:"connectorId" yaml:"connectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#description MagicTransitSite#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Site high availability mode.
	//
	// If set to true, the site can have two connectors and runs in high availability mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#ha_mode MagicTransitSite#ha_mode}
	HaMode interface{} `field:"optional" json:"haMode" yaml:"haMode"`
	// Location of site in latitude and longitude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#location MagicTransitSite#location}
	Location *MagicTransitSiteLocation `field:"optional" json:"location" yaml:"location"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/magic_transit_site#secondary_connector_id MagicTransitSite#secondary_connector_id}
	SecondaryConnectorId *string `field:"optional" json:"secondaryConnectorId" yaml:"secondaryConnectorId"`
}

