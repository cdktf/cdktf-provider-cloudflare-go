package ruleset


type RulesetRulesActionParametersCacheKey struct {
	// Cache by device type. Conflicts with "custom_key.user.device_type".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cache_by_device_type Ruleset#cache_by_device_type}
	CacheByDeviceType interface{} `field:"optional" json:"cacheByDeviceType" yaml:"cacheByDeviceType"`
	// Cache deception armor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cache_deception_armor Ruleset#cache_deception_armor}
	CacheDeceptionArmor interface{} `field:"optional" json:"cacheDeceptionArmor" yaml:"cacheDeceptionArmor"`
	// custom_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#custom_key Ruleset#custom_key}
	CustomKey *RulesetRulesActionParametersCacheKeyCustomKey `field:"optional" json:"customKey" yaml:"customKey"`
	// Ignore query strings order.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#ignore_query_strings_order Ruleset#ignore_query_strings_order}
	IgnoreQueryStringsOrder interface{} `field:"optional" json:"ignoreQueryStringsOrder" yaml:"ignoreQueryStringsOrder"`
}

