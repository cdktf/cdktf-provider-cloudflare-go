package emailroutingcatchall


type EmailRoutingCatchAllAction struct {
	// Type of supported action.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array with items in the following form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/email_routing_catch_all#value EmailRoutingCatchAll#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

