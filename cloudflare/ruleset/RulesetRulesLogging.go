// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesLogging struct {
	// Override the default logging behavior when a rule is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

