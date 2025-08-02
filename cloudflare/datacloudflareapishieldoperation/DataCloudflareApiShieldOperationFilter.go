// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareapishieldoperation


type DataCloudflareApiShieldOperationFilter struct {
	// Direction to order results. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#direction DataCloudflareApiShieldOperation#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter results to only include endpoints containing this pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#endpoint DataCloudflareApiShieldOperation#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Add feature(s) to the results.
	//
	// The feature name that is given here corresponds to the resulting feature object. Have a look at the top-level object description for more details on the specific meaning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#feature DataCloudflareApiShieldOperation#feature}
	Feature *[]*string `field:"optional" json:"feature" yaml:"feature"`
	// Filter results to only include the specified hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#host DataCloudflareApiShieldOperation#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
	// Filter results to only include the specified HTTP methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#method DataCloudflareApiShieldOperation#method}
	Method *[]*string `field:"optional" json:"method" yaml:"method"`
	// Field to order by.
	//
	// When requesting a feature, the feature keys are available for ordering as well, e.g., `thresholds.suggested_threshold`.
	// Available values: "method", "host", "endpoint", "thresholds.$key".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/api_shield_operation#order DataCloudflareApiShieldOperation#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

