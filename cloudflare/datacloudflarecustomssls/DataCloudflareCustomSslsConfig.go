// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomssls

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCustomSslsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_ssls#zone_id DataCloudflareCustomSsls#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Whether to match all search requirements or at least one (any). Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_ssls#match DataCloudflareCustomSsls#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_ssls#max_items DataCloudflareCustomSsls#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Status of the zone's custom SSL. Available values: "active", "expired", "deleted", "pending", "initializing".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/custom_ssls#status DataCloudflareCustomSsls#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

