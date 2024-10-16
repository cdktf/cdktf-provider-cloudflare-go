// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infrastructureaccesstarget


type InfrastructureAccessTargetIp struct {
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/infrastructure_access_target#ipv4 InfrastructureAccessTarget#ipv4}
	Ipv4 *InfrastructureAccessTargetIpIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/infrastructure_access_target#ipv6 InfrastructureAccessTarget#ipv6}
	Ipv6 *InfrastructureAccessTargetIpIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

