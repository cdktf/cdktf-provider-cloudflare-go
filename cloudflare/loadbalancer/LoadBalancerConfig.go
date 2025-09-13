// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerConfig struct {
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
	// A list of pool IDs ordered by their failover priority.
	//
	// Pools defined here are used by default, or when region_pools are not configured for a given region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#default_pools LoadBalancer#default_pools}
	DefaultPools *[]*string `field:"required" json:"defaultPools" yaml:"defaultPools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#fallback_pool LoadBalancer#fallback_pool}
	FallbackPool *string `field:"required" json:"fallbackPool" yaml:"fallbackPool"`
	// The DNS hostname to associate with your Load Balancer.
	//
	// If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#name LoadBalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#zone_id LoadBalancer#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests.
	//
	// For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#adaptive_routing LoadBalancer#adaptive_routing}
	AdaptiveRouting *LoadBalancerAdaptiveRouting `field:"optional" json:"adaptiveRouting" yaml:"adaptiveRouting"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country.
	//
	// Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#country_pools LoadBalancer#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// Object description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#description LoadBalancer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable (the default) this load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#enabled LoadBalancer#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#location_strategy LoadBalancer#location_strategy}
	LocationStrategy *LoadBalancerLocationStrategy `field:"optional" json:"locationStrategy" yaml:"locationStrategy"`
	// List of networks where Load Balancer or Pool is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#networks LoadBalancer#networks}
	Networks *[]*string `field:"optional" json:"networks" yaml:"networks"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter).
	//
	// Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#pop_pools LoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#proxied LoadBalancer#proxied}
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// Configures pool weights.
	//
	// - `steering_policy="random"`: A random pool is selected with probability proportional to pool weights.
	// - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each pool's outstanding requests.
	// - `steering_policy="least_connections"`: Use pool weights to scale each pool's open connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#random_steering LoadBalancer#random_steering}
	RandomSteering *LoadBalancerRandomSteering `field:"optional" json:"randomSteering" yaml:"randomSteering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region.
	//
	// Any regions not explicitly defined will fall back to using default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#region_pools LoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// BETA Field Not General Access: A list of rules for this load balancer to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#rules LoadBalancer#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Specifies the type of session affinity the load balancer should use unless specified as `"none"`.
	//
	// The supported types are: - `"cookie"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin selection is stable and based on the client's ip address. - `"header"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.
	// Available values: "none", "cookie", "ip_cookie", "header".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#session_affinity LoadBalancer#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Configures attributes for session affinity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}
	SessionAffinityAttributes *LoadBalancerSessionAffinityAttributes `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// Time, in seconds, until a client's session expires after being created.
	//
	// Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are: - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800]. - `"header"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}
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
	// Available values: "off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#steering_policy LoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
	//
	// This only applies to gray-clouded (unproxied) load balancers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#ttl LoadBalancer#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

