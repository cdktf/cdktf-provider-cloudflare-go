// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled.
	//
	// If this is true and a request meets eligibility criteria, Cloudflare will write the resource to cache reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/ruleset#eligible Ruleset#eligible}
	Eligible interface{} `field:"required" json:"eligible" yaml:"eligible"`
	// The minimum file size eligible for store in cache reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/ruleset#minimum_file_size Ruleset#minimum_file_size}
	MinimumFileSize *float64 `field:"required" json:"minimumFileSize" yaml:"minimumFileSize"`
}

