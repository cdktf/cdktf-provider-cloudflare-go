// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#cookie Ruleset#cookie}
	Cookie *RulesetRulesActionParametersCacheKeyCustomKeyCookie `field:"optional" json:"cookie" yaml:"cookie"`
	// The header names and values to include in building the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#header Ruleset#header}
	Header *RulesetRulesActionParametersCacheKeyCustomKeyHeader `field:"optional" json:"header" yaml:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#host Ruleset#host}
	Host *RulesetRulesActionParametersCacheKeyCustomKeyHost `field:"optional" json:"host" yaml:"host"`
	// Use the presence of parameters in the query string to build the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#query_string Ruleset#query_string}
	QueryString *RulesetRulesActionParametersCacheKeyCustomKeyQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// Characteristics of the request user agent used in building the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#user Ruleset#user}
	User *RulesetRulesActionParametersCacheKeyCustomKeyUser `field:"optional" json:"user" yaml:"user"`
}

