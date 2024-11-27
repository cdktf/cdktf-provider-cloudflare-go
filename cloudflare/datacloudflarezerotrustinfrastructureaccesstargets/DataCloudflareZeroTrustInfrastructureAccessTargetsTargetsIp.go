// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustinfrastructureaccesstargets


type DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIp struct {
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/data-sources/zero_trust_infrastructure_access_targets#ipv4 DataCloudflareZeroTrustInfrastructureAccessTargets#ipv4}
	Ipv4 *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/data-sources/zero_trust_infrastructure_access_targets#ipv6 DataCloudflareZeroTrustInfrastructureAccessTargets#ipv6}
	Ipv6 *DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

