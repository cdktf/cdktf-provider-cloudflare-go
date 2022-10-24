package loadbalancer


type LoadBalancerRulesOverridesCountryPools struct {
	// See [`country`](#country).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#country LoadBalancer#country}
	Country *string `field:"required" json:"country" yaml:"country"`
	// See [`pool_ids`](#pool_ids).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pool_ids LoadBalancer#pool_ids}
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
}

