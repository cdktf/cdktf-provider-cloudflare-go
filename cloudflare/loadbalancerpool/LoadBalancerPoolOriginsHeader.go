// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolOriginsHeader struct {
	// HTTP Header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/load_balancer_pool#header LoadBalancerPool#header}
	Header *string `field:"required" json:"header" yaml:"header"`
	// Values for the HTTP headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/load_balancer_pool#values LoadBalancerPool#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

