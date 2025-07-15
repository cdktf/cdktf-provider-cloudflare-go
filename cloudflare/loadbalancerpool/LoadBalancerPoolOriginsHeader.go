// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	//
	// Current support is 1 'Host' header override per origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/load_balancer_pool#host LoadBalancerPool#host}
	Host *[]*string `field:"optional" json:"host" yaml:"host"`
}

