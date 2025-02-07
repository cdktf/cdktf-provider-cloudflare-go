// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersBrowserTtl struct {
	// Determines which browser ttl mode to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The TTL (in seconds) if you choose override_origin mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
}

