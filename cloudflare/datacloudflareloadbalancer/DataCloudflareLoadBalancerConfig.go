// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareloadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareLoadBalancerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#zone_id DataCloudflareLoadBalancer#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country.
	//
	// Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#country_pools DataCloudflareLoadBalancer#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#load_balancer_id DataCloudflareLoadBalancer#load_balancer_id}.
	LoadBalancerId *string `field:"optional" json:"loadBalancerId" yaml:"loadBalancerId"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter).
	//
	// Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#pop_pools DataCloudflareLoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region.
	//
	// Any regions not explicitly defined will fall back to using default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#region_pools DataCloudflareLoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// Time, in seconds, until a client's session expires after being created.
	//
	// Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are:
	// - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800].
	// - `"header"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#session_affinity_ttl DataCloudflareLoadBalancer#session_affinity_ttl}
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// Steering Policy for this load balancer.
	//
	// - `"off"`: Use `default_pools`.
	// - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.
	// - `"random"`: Select a pool randomly.
	// - `"dynamic_latency"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).
	// - `"proximity"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.
	// - `"least_outstanding_requests"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.
	// - `"least_connections"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
	// - `""`: Will map to `"geo"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#steering_policy DataCloudflareLoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
	//
	// This only applies to gray-clouded (unproxied) load balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/load_balancer#ttl DataCloudflareLoadBalancer#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

