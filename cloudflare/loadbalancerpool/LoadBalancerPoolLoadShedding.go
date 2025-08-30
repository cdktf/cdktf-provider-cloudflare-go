// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	//
	// Applies to new sessions and traffic without session affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer_pool#default_percent LoadBalancerPool#default_percent}
	DefaultPercent *float64 `field:"optional" json:"defaultPercent" yaml:"defaultPercent"`
	// The default policy to use when load shedding.
	//
	// A random policy randomly sheds a given percent of requests. A hash policy computes a hash over the CF-Connecting-IP address and sheds all requests originating from a percent of IPs.
	// Available values: "random", "hash".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer_pool#default_policy LoadBalancerPool#default_policy}
	DefaultPolicy *string `field:"optional" json:"defaultPolicy" yaml:"defaultPolicy"`
	// The percent of existing sessions to shed from the pool, according to the session policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer_pool#session_percent LoadBalancerPool#session_percent}
	SessionPercent *float64 `field:"optional" json:"sessionPercent" yaml:"sessionPercent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential decay). Available values: "hash".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer_pool#session_policy LoadBalancerPool#session_policy}
	SessionPolicy *string `field:"optional" json:"sessionPolicy" yaml:"sessionPolicy"`
}

