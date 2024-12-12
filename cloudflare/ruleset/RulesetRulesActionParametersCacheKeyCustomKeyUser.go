// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyUser struct {
	// Add device type to the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/ruleset#device_type Ruleset#device_type}
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Add geo data to the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/ruleset#geo Ruleset#geo}
	Geo interface{} `field:"optional" json:"geo" yaml:"geo"`
	// Add language data to the custom key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/ruleset#lang Ruleset#lang}
	Lang interface{} `field:"optional" json:"lang" yaml:"lang"`
}

