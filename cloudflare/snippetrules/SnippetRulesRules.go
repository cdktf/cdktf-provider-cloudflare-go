// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippetrules


type SnippetRulesRules struct {
	// Criteria for an HTTP request to trigger the snippet rule.
	//
	// Uses the Firewall Rules expression language based on Wireshark display filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/snippet_rules#expression SnippetRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Name of the snippet invoked by this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/snippet_rules#snippet_name SnippetRules#snippet_name}
	SnippetName *string `field:"required" json:"snippetName" yaml:"snippetName"`
	// Brief summary of the snippet rule and its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/snippet_rules#description SnippetRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the headers rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/snippet_rules#enabled SnippetRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

