package emailroutingcatchall


type EmailRoutingCatchAllMatcher struct {
	// Type of matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Field for type matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#field EmailRoutingCatchAll#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Value for matcher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#value EmailRoutingCatchAll#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

