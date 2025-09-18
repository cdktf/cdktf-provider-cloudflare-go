// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tieredcache

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TieredCacheConfig struct {
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
	// Enable or disable the Smart Tiered Cache. Available values: "on", "off".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/tiered_cache#value TieredCache#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/tiered_cache#zone_id TieredCache#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

