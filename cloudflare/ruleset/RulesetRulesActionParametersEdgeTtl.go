package ruleset


type RulesetRulesActionParametersEdgeTtl struct {
	// Mode of the edge TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.0/docs/resources/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Default edge TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.0/docs/resources/ruleset#default Ruleset#default}
	Default *float64 `field:"optional" json:"default" yaml:"default"`
	// status_code_ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.0/docs/resources/ruleset#status_code_ttl Ruleset#status_code_ttl}
	StatusCodeTtl interface{} `field:"optional" json:"statusCodeTtl" yaml:"statusCodeTtl"`
}

