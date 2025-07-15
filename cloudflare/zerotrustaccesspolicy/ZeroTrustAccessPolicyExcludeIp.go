// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeIp struct {
	// An IPv4 or IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_policy#ip ZeroTrustAccessPolicy#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

