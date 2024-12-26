// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolOriginSteering struct {
	// Origin steering policy to be used.
	//
	// Value `random` selects an origin randomly. Value `hash` selects an origin by computing a hash over the CF-Connecting-IP address. Value `least_outstanding_requests` selects an origin by taking into consideration origin weights, as well as each origin's number of outstanding requests. Origins with more pending requests are weighted proportionately less relative to others. Value `least_connections` selects an origin by taking into consideration origin weights, as well as each origin's number of open connections. Origins with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections. Available values: `""`, `hash`, `random`, `least_outstanding_requests`, `least_connections`. Defaults to `random`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/load_balancer_pool#policy LoadBalancerPool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

