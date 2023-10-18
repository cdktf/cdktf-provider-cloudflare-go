// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolOriginSteering struct {
	// Origin steering policy to be used.
	//
	// Value `random` selects an origin randomly. Value `hash` selects an origin by computing a hash over the CF-Connecting-IP address. Value `least_outstanding_requests` selects an origin by taking into consideration origin weights, as well as each origin's number of outstanding requests. Origins with more pending requests are weighted proportionately less relative to others. Available values: `""`, `hash`, `random`, `least_outstanding_requests`. Defaults to `random`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.17.0/docs/resources/load_balancer_pool#policy LoadBalancerPool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

