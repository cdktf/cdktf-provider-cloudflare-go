// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareapishieldschemas

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareApiShieldSchemasConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/api_shield_schemas#zone_id DataCloudflareApiShieldSchemas#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/api_shield_schemas#max_items DataCloudflareApiShieldSchemas#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/api_shield_schemas#omit_source DataCloudflareApiShieldSchemas#omit_source}
	OmitSource interface{} `field:"optional" json:"omitSource" yaml:"omitSource"`
	// Flag whether schema is enabled for validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/api_shield_schemas#validation_enabled DataCloudflareApiShieldSchemas#validation_enabled}
	ValidationEnabled interface{} `field:"optional" json:"validationEnabled" yaml:"validationEnabled"`
}

