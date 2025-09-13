// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersEdgeTtlStatusCodeTtl struct {
	// The time to cache the response for (in seconds).
	//
	// A value of 0 is equivalent to setting the cache control header with the value "no-cache". A value of -1 is equivalent to setting the cache control header with the value of "no-store".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#value Ruleset#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// A single status code to apply the TTL to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// A range of status codes to apply the TTL to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#status_code_range Ruleset#status_code_range}
	StatusCodeRange *RulesetRulesActionParametersEdgeTtlStatusCodeTtlStatusCodeRange `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
}

