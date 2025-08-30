// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule and category overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#categories Ruleset#categories}
	Categories interface{} `field:"optional" json:"categories" yaml:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than rule and category overrides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#rules Ruleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// A sensitivity level to set for all rules.
	//
	// This option has lower precedence than rule and category overrides and is only applicable for DDoS phases.
	// Available values: "default", "medium", "low", "eoff".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#sensitivity_level Ruleset#sensitivity_level}
	SensitivityLevel *string `field:"optional" json:"sensitivityLevel" yaml:"sensitivityLevel"`
}

