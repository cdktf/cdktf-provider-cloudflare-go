// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezonelockdowns

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZoneLockdownsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#zone_id DataCloudflareZoneLockdowns#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// The timestamp of when the rule was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#created_on DataCloudflareZoneLockdowns#created_on}
	CreatedOn *string `field:"optional" json:"createdOn" yaml:"createdOn"`
	// A string to search for in the description of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#description DataCloudflareZoneLockdowns#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string to search for in the description of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#description_search DataCloudflareZoneLockdowns#description_search}
	DescriptionSearch *string `field:"optional" json:"descriptionSearch" yaml:"descriptionSearch"`
	// A single IP address to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#ip DataCloudflareZoneLockdowns#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// A single IP address range to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#ip_range_search DataCloudflareZoneLockdowns#ip_range_search}
	IpRangeSearch *string `field:"optional" json:"ipRangeSearch" yaml:"ipRangeSearch"`
	// A single IP address to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#ip_search DataCloudflareZoneLockdowns#ip_search}
	IpSearch *string `field:"optional" json:"ipSearch" yaml:"ipSearch"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#max_items DataCloudflareZoneLockdowns#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The timestamp of when the rule was last modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#modified_on DataCloudflareZoneLockdowns#modified_on}
	ModifiedOn *string `field:"optional" json:"modifiedOn" yaml:"modifiedOn"`
	// The priority of the rule to control the processing order.
	//
	// A lower number indicates higher priority. If not provided, any rules with a configured priority will be processed before rules without a priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#priority DataCloudflareZoneLockdowns#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/zone_lockdowns#uri_search DataCloudflareZoneLockdowns#uri_search}
	UriSearch *string `field:"optional" json:"uriSearch" yaml:"uriSearch"`
}

