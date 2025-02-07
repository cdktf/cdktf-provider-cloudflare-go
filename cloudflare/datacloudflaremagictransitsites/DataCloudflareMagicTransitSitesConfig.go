// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremagictransitsites

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareMagicTransitSitesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/magic_transit_sites#account_id DataCloudflareMagicTransitSites#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/magic_transit_sites#connector_identifier DataCloudflareMagicTransitSites#connector_identifier}
	ConnectorIdentifier *string `field:"optional" json:"connectorIdentifier" yaml:"connectorIdentifier"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/magic_transit_sites#max_items DataCloudflareMagicTransitSites#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
}

