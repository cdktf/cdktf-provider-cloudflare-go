// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheReserve struct {
	// Whether Cache Reserve is enabled.
	//
	// If this is true and a request meets eligibility criteria, Cloudflare will write the resource to Cache Reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#eligible Ruleset#eligible}
	Eligible interface{} `field:"required" json:"eligible" yaml:"eligible"`
	// The minimum file size eligible for storage in Cache Reserve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#minimum_file_size Ruleset#minimum_file_size}
	MinimumFileSize *float64 `field:"optional" json:"minimumFileSize" yaml:"minimumFileSize"`
}

