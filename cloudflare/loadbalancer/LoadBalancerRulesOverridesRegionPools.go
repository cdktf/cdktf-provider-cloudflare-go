// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerRulesOverridesRegionPools struct {
	// A list of pool IDs in failover priority to use in the given region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.34.0/docs/resources/load_balancer#pool_ids LoadBalancer#pool_ids}
	PoolIds *[]*string `field:"required" json:"poolIds" yaml:"poolIds"`
	// A region code which must be in the list defined [here](https://developers.cloudflare.com/load-balancing/reference/region-mapping-api/#list-of-load-balancer-regions). Multiple entries should not be specified with the same region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.34.0/docs/resources/load_balancer#region LoadBalancer#region}
	Region *string `field:"required" json:"region" yaml:"region"`
}

