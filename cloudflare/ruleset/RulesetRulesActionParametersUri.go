// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersUri struct {
	// Path portion rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/ruleset#path Ruleset#path}
	Path *RulesetRulesActionParametersUriPath `field:"optional" json:"path" yaml:"path"`
	// Query portion rewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/ruleset#query Ruleset#query}
	Query *RulesetRulesActionParametersUriQuery `field:"optional" json:"query" yaml:"query"`
}

