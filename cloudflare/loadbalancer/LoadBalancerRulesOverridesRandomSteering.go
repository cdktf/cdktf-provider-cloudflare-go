package loadbalancer


type LoadBalancerRulesOverridesRandomSteering struct {
	// See [`default_weight`](#default_weight).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#default_weight LoadBalancer#default_weight}
	DefaultWeight *float64 `field:"optional" json:"defaultWeight" yaml:"defaultWeight"`
	// See [`pool_weights`](#pool_weights).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pool_weights LoadBalancer#pool_weights}
	PoolWeights *map[string]*float64 `field:"optional" json:"poolWeights" yaml:"poolWeights"`
}

