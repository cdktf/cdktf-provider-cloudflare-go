// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerpool


type LoadBalancerPoolOrigins struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable hostname.
	//
	// Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by Cloudflare. To set an internal/reserved address, virtual_network_id must also be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#address LoadBalancerPool#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Whether to enable (the default) this origin within the pool.
	//
	// Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#enabled LoadBalancerPool#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The request header is used to pass additional information with an HTTP request. Currently supported header is 'Host'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#header LoadBalancerPool#header}
	Header *LoadBalancerPoolOriginsHeader `field:"optional" json:"header" yaml:"header"`
	// A human-identifiable name for the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#name LoadBalancerPool#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port for upstream connections. A value of 0 means the default port for the protocol will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#port LoadBalancerPool#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also belong to the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#virtual_network_id LoadBalancerPool#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// The weight of this origin relative to other origins in the pool.
	//
	// Based on the configured weight the total traffic is distributed among origins within the pool.
	// - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the origin's outstanding requests.
	// - `origin_steering.policy="least_connections"`: Use weight to scale the origin's open connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/load_balancer_pool#weight LoadBalancerPool#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

