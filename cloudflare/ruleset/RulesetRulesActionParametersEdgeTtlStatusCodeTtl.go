// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtl struct {
	// Status code for which the edge TTL is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.16.0/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// status_code_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.16.0/docs/resources/ruleset#status_code_range Ruleset#status_code_range}
	StatusCodeRange interface{} `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
	// Status code edge TTL value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.16.0/docs/resources/ruleset#value Ruleset#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

