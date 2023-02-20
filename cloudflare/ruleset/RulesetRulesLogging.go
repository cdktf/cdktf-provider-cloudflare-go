package ruleset


type RulesetRulesLogging struct {
	// Override the default logging behavior when a rule is matched.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#enabled Ruleset#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Override the default logging behavior when a rule is matched. Available values: `enabled`, `disabled`. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status Ruleset#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

