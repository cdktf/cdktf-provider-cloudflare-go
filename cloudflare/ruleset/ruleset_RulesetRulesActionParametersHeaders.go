package ruleset


type RulesetRulesActionParametersHeaders struct {
	// Use a value dynamically determined by the Firewall Rules expression language based on Wireshark display filters.
	//
	// Refer to the [Firewall Rules language](https://developers.cloudflare.com/firewall/cf-firewall-language) documentation for all available fields, operators, and functions. Conflicts with `"value"`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#expression Ruleset#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Name of the HTTP request header to target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#name Ruleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Action to perform on the HTTP request header. Available values: `remove`, `set`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#operation Ruleset#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Static value to provide as the HTTP request header value. Conflicts with `"expression"`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#value Ruleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

