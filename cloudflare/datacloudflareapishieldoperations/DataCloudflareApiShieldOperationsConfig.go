// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareapishieldoperations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareApiShieldOperationsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#zone_id DataCloudflareApiShieldOperations#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#direction DataCloudflareApiShieldOperations#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter results to only include endpoints containing this pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#endpoint DataCloudflareApiShieldOperations#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Add feature(s) to the results.
	//
	// The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#feature DataCloudflareApiShieldOperations#feature}
	Feature *[]*string `field:"optional" json:"feature" yaml:"feature"`
	// Filter results to only include the specified hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#host DataCloudflareApiShieldOperations#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#max_items DataCloudflareApiShieldOperations#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Filter results to only include the specified HTTP methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#method DataCloudflareApiShieldOperations#method}
	Method *[]*string `field:"optional" json:"method" yaml:"method"`
	// Field to order by.
	//
	// When requesting a feature, the feature keys are available for ordering as well, e.g., `thresholds.suggested_threshold`.
	// Available values: "method", "host", "endpoint", "thresholds.$key".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/api_shield_operations#order DataCloudflareApiShieldOperations#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

