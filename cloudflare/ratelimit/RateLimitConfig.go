// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ratelimit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RateLimitConfig struct {
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
	// The action to perform when the threshold of matched traffic within the configured period is exceeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#action RateLimit#action}
	Action *RateLimitAction `field:"required" json:"action" yaml:"action"`
	// Determines which traffic the rate limit counts towards the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#match RateLimit#match}
	Match *RateLimitMatch `field:"required" json:"match" yaml:"match"`
	// The time in seconds (an integer value) to count matching traffic.
	//
	// If the count exceeds the configured threshold within this period, Cloudflare will perform the configured action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#period RateLimit#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The threshold that will trigger the configured mitigation action.
	//
	// Configure this value along with the `period` property to establish a threshold per period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#threshold RateLimit#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/rate_limit#zone_id RateLimit#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

