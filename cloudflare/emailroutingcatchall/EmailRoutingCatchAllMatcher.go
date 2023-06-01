package emailroutingcatchall


type EmailRoutingCatchAllMatcher struct {
	// Type of matcher. Available values: `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.7.1/docs/resources/email_routing_catch_all#type EmailRoutingCatchAll#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

