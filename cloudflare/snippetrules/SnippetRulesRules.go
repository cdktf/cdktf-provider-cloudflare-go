// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippetrules


type SnippetRulesRules struct {
	// The expression defining which traffic will match the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/snippet_rules#expression SnippetRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The identifying name of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/snippet_rules#snippet_name SnippetRules#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// An informative description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/snippet_rules#description SnippetRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the rule should be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/snippet_rules#enabled SnippetRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

