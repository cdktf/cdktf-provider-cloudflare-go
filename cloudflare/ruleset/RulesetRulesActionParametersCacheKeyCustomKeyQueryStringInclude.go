// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Whether to include all query string parameters in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#all Ruleset#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// A list of query string parameters to include in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#list Ruleset#list}
	List *[]*string `field:"optional" json:"list" yaml:"list"`
}

