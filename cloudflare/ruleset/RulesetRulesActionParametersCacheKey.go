// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersCacheKey struct {
	// Cache by device type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/ruleset#cache_by_device_type Ruleset#cache_by_device_type}
	CacheByDeviceType interface{} `field:"optional" json:"cacheByDeviceType" yaml:"cacheByDeviceType"`
	// Cache deception armor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/ruleset#cache_deception_armor Ruleset#cache_deception_armor}
	CacheDeceptionArmor interface{} `field:"optional" json:"cacheDeceptionArmor" yaml:"cacheDeceptionArmor"`
	// custom_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/ruleset#custom_key Ruleset#custom_key}
	CustomKey interface{} `field:"optional" json:"customKey" yaml:"customKey"`
	// Ignore query strings order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/ruleset#ignore_query_strings_order Ruleset#ignore_query_strings_order}
	IgnoreQueryStringsOrder interface{} `field:"optional" json:"ignoreQueryStringsOrder" yaml:"ignoreQueryStringsOrder"`
}

