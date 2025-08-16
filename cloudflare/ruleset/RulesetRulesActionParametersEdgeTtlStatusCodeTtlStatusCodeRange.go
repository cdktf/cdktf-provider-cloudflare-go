// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtlStatusCodeRange struct {
	// The lower bound of the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/ruleset#from Ruleset#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The upper bound of the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/ruleset#to Ruleset#to}
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

