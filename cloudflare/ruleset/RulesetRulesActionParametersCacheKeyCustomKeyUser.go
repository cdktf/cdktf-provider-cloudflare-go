// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/ruleset#device_type Ruleset#device_type}
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Use the user agents's country in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/ruleset#geo Ruleset#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// Use the user agent's language in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/ruleset#lang Ruleset#lang}
	Lang interface{} `field:"optional" json:"lang" yaml:"lang"`
}

