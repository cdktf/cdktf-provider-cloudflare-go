// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKey struct {
	// Which cookies to include in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#cookie Ruleset#cookie}
	Cookie *RulesetRulesActionParametersCacheKeyCustomKeyCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// Which headers to include in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#header Ruleset#header}
	Header *RulesetRulesActionParametersCacheKeyCustomKeyHeader `field:"optional" json:"header" yaml:"header"`
	// How to use the host in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#host Ruleset#host}
	Host *RulesetRulesActionParametersCacheKeyCustomKeyHost `field:"optional" json:"host" yaml:"host"`
	// Which query string parameters to include in or exclude from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#query_string Ruleset#query_string}
	QueryString *RulesetRulesActionParametersCacheKeyCustomKeyQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// How to use characteristics of the request user agent in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#user Ruleset#user}
	User *RulesetRulesActionParametersCacheKeyCustomKeyUser `field:"optional" json:"user" yaml:"user"`
}

