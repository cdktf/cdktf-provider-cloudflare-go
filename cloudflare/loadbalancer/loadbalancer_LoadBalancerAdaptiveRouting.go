package loadbalancer


type LoadBalancerAdaptiveRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#failover_across_pools LoadBalancer#failover_across_pools}.
	FailoverAcrossPools interface{} `field:"optional" json:"failoverAcrossPools" yaml:"failoverAcrossPools"`
}

