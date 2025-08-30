// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessinfrastructuretarget


type ZeroTrustAccessInfrastructureTargetIp struct {
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_infrastructure_target#ipv4 ZeroTrustAccessInfrastructureTarget#ipv4}
	Ipv4 *ZeroTrustAccessInfrastructureTargetIpIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_infrastructure_target#ipv6 ZeroTrustAccessInfrastructureTarget#ipv6}
	Ipv6 *ZeroTrustAccessInfrastructureTargetIpIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

