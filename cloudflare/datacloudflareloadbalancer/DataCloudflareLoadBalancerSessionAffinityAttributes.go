// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareloadbalancer


type DataCloudflareLoadBalancerSessionAffinityAttributes struct {
	// Configures the drain duration in seconds.
	//
	// This field is only used when session affinity is enabled on the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/load_balancer#drain_duration DataCloudflareLoadBalancer#drain_duration}
	DrainDuration *float64 `field:"optional" json:"drainDuration" yaml:"drainDuration"`
}

