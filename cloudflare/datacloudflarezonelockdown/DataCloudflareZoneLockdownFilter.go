// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezonelockdown


type DataCloudflareZoneLockdownFilter struct {
	// The timestamp of when the rule was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#created_on DataCloudflareZoneLockdown#created_on}
	CreatedOn *string `field:"optional" json:"createdOn" yaml:"createdOn"`
	// A string to search for in the description of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#description DataCloudflareZoneLockdown#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string to search for in the description of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#description_search DataCloudflareZoneLockdown#description_search}
	DescriptionSearch *string `field:"optional" json:"descriptionSearch" yaml:"descriptionSearch"`
	// A single IP address to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#ip DataCloudflareZoneLockdown#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// A single IP address range to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#ip_range_search DataCloudflareZoneLockdown#ip_range_search}
	IpRangeSearch *string `field:"optional" json:"ipRangeSearch" yaml:"ipRangeSearch"`
	// A single IP address to search for in existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#ip_search DataCloudflareZoneLockdown#ip_search}
	IpSearch *string `field:"optional" json:"ipSearch" yaml:"ipSearch"`
	// The timestamp of when the rule was last modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#modified_on DataCloudflareZoneLockdown#modified_on}
	ModifiedOn *string `field:"optional" json:"modifiedOn" yaml:"modifiedOn"`
	// The priority of the rule to control the processing order.
	//
	// A lower number indicates higher priority. If not provided, any rules with a configured priority will be processed before rules without a priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#priority DataCloudflareZoneLockdown#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/zone_lockdown#uri_search DataCloudflareZoneLockdown#uri_search}
	UriSearch *string `field:"optional" json:"uriSearch" yaml:"uriSearch"`
}

