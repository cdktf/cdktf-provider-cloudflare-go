// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersSni struct {
	// Value to define for SNI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.23.0/docs/resources/ruleset#value Ruleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

