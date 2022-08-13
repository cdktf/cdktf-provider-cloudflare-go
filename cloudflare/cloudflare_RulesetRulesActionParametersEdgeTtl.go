// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RulesetRulesActionParametersEdgeTtl struct {
	// Default edge TTL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#default Ruleset#default}
	Default *float64 `field:"required" json:"default" yaml:"default"`
	// Mode of the edge TTL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#mode Ruleset#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// status_code_ttl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code_ttl Ruleset#status_code_ttl}
	StatusCodeTtl interface{} `field:"optional" json:"statusCodeTtl" yaml:"statusCodeTtl"`
}

