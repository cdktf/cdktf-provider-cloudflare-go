package loadbalancer


type LoadBalancerRulesOverridesLocationStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#mode LoadBalancer#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#prefer_ecs LoadBalancer#prefer_ecs}.
	PreferEcs *string `field:"optional" json:"preferEcs" yaml:"preferEcs"`
}

