// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustinfrastructureaccesstargets


type DataCloudflareZeroTrustInfrastructureAccessTargetsTargets struct {
	// The IPv4/IPv6 address that identifies where to reach a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/data-sources/zero_trust_infrastructure_access_targets#ip DataCloudflareZeroTrustInfrastructureAccessTargets#ip}
	Ip *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp `field:"required" json:"ip" yaml:"ip"`
}

