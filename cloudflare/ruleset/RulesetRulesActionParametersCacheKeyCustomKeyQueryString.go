// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyQueryString struct {
	// Which query string parameters to exclude from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#exclude Ruleset#exclude}
	Exclude *RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// Which query string parameters to include in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#include Ruleset#include}
	Include *RulesetRulesActionParametersCacheKeyCustomKeyQueryStringInclude `field:"optional" json:"include" yaml:"include"`
}

