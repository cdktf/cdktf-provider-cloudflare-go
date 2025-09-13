// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKey struct {
	// Whether to separate cached content based on the visitor's device type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#cache_by_device_type Ruleset#cache_by_device_type}
	CacheByDeviceType interface{} `field:"optional" json:"cacheByDeviceType" yaml:"cacheByDeviceType"`
	// Whether to protect from web cache deception attacks, while allowing static assets to be cached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#cache_deception_armor Ruleset#cache_deception_armor}
	CacheDeceptionArmor interface{} `field:"optional" json:"cacheDeceptionArmor" yaml:"cacheDeceptionArmor"`
	// Which components of the request are included or excluded from the cache key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#custom_key Ruleset#custom_key}
	CustomKey *RulesetRulesActionParametersCacheKeyCustomKey `field:"optional" json:"customKey" yaml:"customKey"`
	// Whether to treat requests with the same query parameters the same, regardless of the order those query parameters are in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/ruleset#ignore_query_strings_order Ruleset#ignore_query_strings_order}
	IgnoreQueryStringsOrder interface{} `field:"optional" json:"ignoreQueryStringsOrder" yaml:"ignoreQueryStringsOrder"`
}

