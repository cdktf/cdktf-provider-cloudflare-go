// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKeyCustomKeyHeader struct {
	// A list of headers to check for the presence of.
	//
	// The presence of these headers is included in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#check_presence Ruleset#check_presence}
	CheckPresence *[]*string `field:"optional" json:"checkPresence" yaml:"checkPresence"`
	// A mapping of header names to a list of values.
	//
	// If a header is present in the request and contains any of the values provided, its value is included in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#contains Ruleset#contains}
	Contains interface{} `field:"optional" json:"contains" yaml:"contains"`
	// Whether to exclude the origin header in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#exclude_origin Ruleset#exclude_origin}
	ExcludeOrigin interface{} `field:"optional" json:"excludeOrigin" yaml:"excludeOrigin"`
	// A list of headers to include in the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/ruleset#include Ruleset#include}
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

