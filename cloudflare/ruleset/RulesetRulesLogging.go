// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesLogging struct {
	// Whether to generate a log when the rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

