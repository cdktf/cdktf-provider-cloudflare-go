// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package regionaltieredcache

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegionalTieredCacheConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/regional_tiered_cache#zone_id RegionalTieredCache#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Value of the Regional Tiered Cache zone setting. Available values: "on", "off".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/regional_tiered_cache#value RegionalTieredCache#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

