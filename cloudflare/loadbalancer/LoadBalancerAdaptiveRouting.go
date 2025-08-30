// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering.
	//
	// When set false (the default) zero-downtime failover will only occur between origins within the same pool. See `session_affinity_attributes` for control over when sessions are broken or reassigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer#failover_across_pools LoadBalancer#failover_across_pools}
	FailoverAcrossPools interface{} `field:"optional" json:"failoverAcrossPools" yaml:"failoverAcrossPools"`
}

