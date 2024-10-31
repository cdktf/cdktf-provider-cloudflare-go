// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheReserve struct {
	// Determines whether Cloudflare will write the eligible resource to cache reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/ruleset#eligible Ruleset#eligible}
	Eligible interface{} `field:"required" json:"eligible" yaml:"eligible"`
	// The minimum file size, in bytes, eligible for storage in cache reserve.
	//
	// If omitted and "eligible" is true, Cloudflare will use 0 bytes by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/ruleset#minimum_file_size Ruleset#minimum_file_size}
	MinimumFileSize *float64 `field:"optional" json:"minimumFileSize" yaml:"minimumFileSize"`
}

