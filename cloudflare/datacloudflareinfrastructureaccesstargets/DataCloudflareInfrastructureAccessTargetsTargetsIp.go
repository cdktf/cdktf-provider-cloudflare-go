// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareinfrastructureaccesstargets


type DataCloudflareInfrastructureAccessTargetsTargetsIp struct {
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/data-sources/infrastructure_access_targets#ipv4 DataCloudflareInfrastructureAccessTargets#ipv4}
	Ipv4 *DataCloudflareInfrastructureAccessTargetsTargetsIpIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/data-sources/infrastructure_access_targets#ipv6 DataCloudflareInfrastructureAccessTargets#ipv6}
	Ipv6 *DataCloudflareInfrastructureAccessTargetsTargetsIpIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

