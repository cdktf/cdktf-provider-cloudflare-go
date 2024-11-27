// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareinfrastructureaccesstargets


type DataCloudflareInfrastructureAccessTargetsTargets struct {
	// The IPv4/IPv6 address that identifies where to reach a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/data-sources/infrastructure_access_targets#ip DataCloudflareInfrastructureAccessTargets#ip}
	Ip *DataCloudflareInfrastructureAccessTargetsTargetsIp `field:"required" json:"ip" yaml:"ip"`
}

