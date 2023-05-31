package ruleset


type RulesetRulesActionParametersBrowserTtl struct {
	// Mode of the browser TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Default browser TTL. This value is required when override_origin is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
}

