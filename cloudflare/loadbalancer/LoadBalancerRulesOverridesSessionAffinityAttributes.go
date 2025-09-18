// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds.
	//
	// This field is only used when session affinity is enabled on the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#drain_duration LoadBalancer#drain_duration}
	DrainDuration *float64 `field:"optional" json:"drainDuration" yaml:"drainDuration"`
	// Configures the names of HTTP headers to base session affinity on when header `session_affinity` is enabled.
	//
	// At least one HTTP header name must be provided. To specify the exact cookies to be used, include an item in the following format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything after the colon is a comma-separated list of cookie names. Providing only `"cookie"` will result in all cookies being used. The default max number of HTTP header names that can be provided depends on your plan: 5 for Enterprise, 1 for all other plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#headers LoadBalancer#headers}
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify how HTTP headers on load balancing requests will be used.
	//
	// The supported values are: - `"true"`: Load balancing requests must contain *all* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created. - `"false"`: Load balancing requests must contain *at least one* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#require_all_headers LoadBalancer#require_all_headers}
	RequireAllHeaders interface{} `field:"optional" json:"requireAllHeaders" yaml:"requireAllHeaders"`
	// Configures the SameSite attribute on session affinity cookie.
	//
	// Value "Auto" will be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note: when using value "None", the secure attribute can not be set to "Never".
	// Available values: "Auto", "Lax", "None", "Strict".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#samesite LoadBalancer#samesite}
	Samesite *string `field:"optional" json:"samesite" yaml:"samesite"`
	// Configures the Secure attribute on session affinity cookie.
	//
	// Value "Always" indicates the Secure attribute will be set in the Set-Cookie header, "Never" indicates the Secure attribute will not be set, and "Auto" will set the Secure attribute depending if Always Use HTTPS is enabled.
	// Available values: "Auto", "Always", "Never".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#secure LoadBalancer#secure}
	Secure *string `field:"optional" json:"secure" yaml:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session affinity is enabled.
	//
	// This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover takes place for sessions pinned to the origin (default). - `"temporary"`: Traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping. - `"sticky"`: The session affinity cookie is updated and subsequent requests are sent to the new origin. Note: Zero-downtime failover with sticky sessions is currently not supported for session affinity by header.
	// Available values: "none", "temporary", "sticky".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/load_balancer#zero_downtime_failover LoadBalancer#zero_downtime_failover}
	ZeroDowntimeFailover *string `field:"optional" json:"zeroDowntimeFailover" yaml:"zeroDowntimeFailover"`
}

