// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustinfrastructureaccesstarget


type ZeroTrustInfrastructureAccessTargetIp struct {
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_infrastructure_access_target#ipv4 ZeroTrustInfrastructureAccessTarget#ipv4}
	Ipv4 *ZeroTrustInfrastructureAccessTargetIpIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_infrastructure_access_target#ipv6 ZeroTrustInfrastructureAccessTarget#ipv6}
	Ipv6 *ZeroTrustInfrastructureAccessTargetIpIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

