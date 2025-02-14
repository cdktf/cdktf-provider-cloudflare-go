// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippetrules


type SnippetRulesRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/snippet_rules#description SnippetRules#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/snippet_rules#enabled SnippetRules#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/snippet_rules#expression SnippetRules#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Snippet identifying name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/snippet_rules#snippet_name SnippetRules#snippet_name}
	SnippetName *string `field:"optional" json:"snippetName" yaml:"snippetName"`
}

