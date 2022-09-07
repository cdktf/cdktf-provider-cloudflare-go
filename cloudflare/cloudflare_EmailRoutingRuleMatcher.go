// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type EmailRoutingRuleMatcher struct {
	// Type of matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Field for type matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_rule#field EmailRoutingRule#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Value for matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_rule#value EmailRoutingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

