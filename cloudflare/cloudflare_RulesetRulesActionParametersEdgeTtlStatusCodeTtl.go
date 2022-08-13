// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RulesetRulesActionParametersEdgeTtlStatusCodeTtl struct {
	// Status code edge TTL value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#value Ruleset#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Status code for which the edge TTL is applied. Conflicts with "status_code_range".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// status_code_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code_range Ruleset#status_code_range}
	StatusCodeRange interface{} `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
}

