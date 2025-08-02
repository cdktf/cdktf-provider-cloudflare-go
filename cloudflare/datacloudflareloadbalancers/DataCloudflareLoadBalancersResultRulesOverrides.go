// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareloadbalancers


type DataCloudflareLoadBalancersResultRulesOverrides struct {
	// A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country.
	//
	// Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/load_balancers#country_pools DataCloudflareLoadBalancers#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter).
	//
	// Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/load_balancers#pop_pools DataCloudflareLoadBalancers#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region.
	//
	// Any regions not explicitly defined will fall back to using default_pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/data-sources/load_balancers#region_pools DataCloudflareLoadBalancers#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
}

