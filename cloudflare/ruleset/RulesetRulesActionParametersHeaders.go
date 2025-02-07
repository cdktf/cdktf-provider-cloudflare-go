// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersHeaders struct {
	// Expression for the header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/ruleset#operation Ruleset#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Static value for the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

