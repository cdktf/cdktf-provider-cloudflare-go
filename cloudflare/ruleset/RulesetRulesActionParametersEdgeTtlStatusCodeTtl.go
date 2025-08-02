// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtl struct {
	// Time to cache a response (in seconds).
	//
	// A value of 0 is equivalent to setting the Cache-Control header with the value "no-cache". A value of -1 is equivalent to setting Cache-Control header with the value of "no-store".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#value Ruleset#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Set the TTL for responses with this specific status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The range of status codes used to apply the selected mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#status_code_range Ruleset#status_code_range}
	StatusCodeRange *RulesetRulesActionParametersEdgeTtlStatusCodeTtlStatusCodeRange `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
}

