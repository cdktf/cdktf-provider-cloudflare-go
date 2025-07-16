// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezones

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZonesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#account DataCloudflareZones#account}.
	Account *DataCloudflareZonesAccount `field:"optional" json:"account" yaml:"account"`
	// Direction to order zones. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#direction DataCloudflareZones#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Whether to match all search requirements or at least one (any). Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#match DataCloudflareZones#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#max_items DataCloudflareZones#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// A domain name.
	//
	// Optional filter operators can be provided to extend refine the search:
	//   * `equal` (default)
	//   * `not_equal`
	//   * `starts_with`
	//   * `ends_with`
	//   * `contains`
	//   * `starts_with_case_sensitive`
	//   * `ends_with_case_sensitive`
	//   * `contains_case_sensitive`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#name DataCloudflareZones#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Field to order zones by. Available values: "name", "status", "account.id", "account.name", "plan.id".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#order DataCloudflareZones#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Specify a zone status to filter by. Available values: "initializing", "pending", "active", "moved".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/zones#status DataCloudflareZones#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

