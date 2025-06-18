// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersFromValue struct {
	// Keep the query string of the original request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#preserve_query_string Ruleset#preserve_query_string}
	PreserveQueryString interface{} `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// The status code to be used for the redirect. Available values: 301, 302, 303, 307, 308.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The URL to redirect the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/ruleset#target_url Ruleset#target_url}
	TargetUrl *RulesetRulesActionParametersFromValueTargetUrl `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}

