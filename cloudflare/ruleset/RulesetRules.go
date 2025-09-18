// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRules struct {
	// The action to perform when the rule matches.
	//
	// Available values: "block", "challenge", "compress_response", "ddos_dynamic", "execute", "force_connection_close", "js_challenge", "log", "log_custom_field", "managed_challenge", "redirect", "rewrite", "route", "score", "serve_error", "set_cache_settings", "set_config", "skip".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#action Ruleset#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The expression defining which traffic will match the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The parameters configuring the rule's action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#action_parameters Ruleset#action_parameters}
	ActionParameters *RulesetRulesActionParameters `field:"optional" json:"actionParameters" yaml:"actionParameters"`
	// An informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#description Ruleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule should be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Configuration for exposed credential checking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#exposed_credential_check Ruleset#exposed_credential_check}
	ExposedCredentialCheck *RulesetRulesExposedCredentialCheck `field:"optional" json:"exposedCredentialCheck" yaml:"exposedCredentialCheck"`
	// An object configuring the rule's logging behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#logging Ruleset#logging}
	Logging *RulesetRulesLogging `field:"optional" json:"logging" yaml:"logging"`
	// An object configuring the rule's rate limit behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#ratelimit Ruleset#ratelimit}
	Ratelimit *RulesetRulesRatelimit `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#ref Ruleset#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

