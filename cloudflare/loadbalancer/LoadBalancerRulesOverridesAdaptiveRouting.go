package loadbalancer


type LoadBalancerRulesOverridesAdaptiveRouting struct {
	// See [`failover_across_pools`](#failover_across_pools).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#failover_across_pools LoadBalancer#failover_across_pools}
	FailoverAcrossPools interface{} `field:"optional" json:"failoverAcrossPools" yaml:"failoverAcrossPools"`
}

