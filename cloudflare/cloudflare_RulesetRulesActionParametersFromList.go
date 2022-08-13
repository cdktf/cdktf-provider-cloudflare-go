// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RulesetRulesActionParametersFromList struct {
	// Expression to use for the list lookup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#key Ruleset#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the list.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

