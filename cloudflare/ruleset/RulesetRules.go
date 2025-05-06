// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRules struct {
	// The action to perform when the rule matches. Available values: "block".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The parameters configuring the rule's action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#action_parameters Ruleset#action_parameters}
	ActionParameters *RulesetRulesActionParameters `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// The categories of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#categories Ruleset#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// An informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#description Ruleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule should be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Configure checks for exposed credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#exposed_credential_check Ruleset#exposed_credential_check}
	ExposedCredentialCheck *RulesetRulesExposedCredentialCheck `field:"optional" json:"exposedCredentialCheck" yaml:"exposedCredentialCheck"`
	// The expression defining which traffic will match the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// An object configuring the rule's logging behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#logging Ruleset#logging}
	Logging *RulesetRulesLogging `field:"optional" json:"logging" yaml:"logging"`
	// An object configuring the rule's ratelimit behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#ratelimit Ruleset#ratelimit}
	Ratelimit *RulesetRulesRatelimit `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#ref Ruleset#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

