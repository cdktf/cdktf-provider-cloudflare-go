// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/load_balancer_pool#origin LoadBalancerPool#origin}
	Origin *LoadBalancerPoolNotificationFilterOrigin `field:"optional" json:"origin" yaml:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to reset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/load_balancer_pool#pool LoadBalancerPool#pool}
	Pool *LoadBalancerPoolNotificationFilterPool `field:"optional" json:"pool" yaml:"pool"`
}

