// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.
	//
	// - `"pop"`: Use the Cloudflare PoP location.
	// - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.
	// Available values: "pop", "resolver_ip".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/load_balancer#mode LoadBalancer#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	// Available values: "always", "never", "proximity", "geo".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/load_balancer#prefer_ecs LoadBalancer#prefer_ecs}
	PreferEcs *string `field:"optional" json:"preferEcs" yaml:"preferEcs"`
}

