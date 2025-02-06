// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationEndpoints struct {
	// doh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_dns_location#doh ZeroTrustDnsLocation#doh}
	Doh *ZeroTrustDnsLocationEndpointsDoh `field:"optional" json:"doh" yaml:"doh"`
	// dot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_dns_location#dot ZeroTrustDnsLocation#dot}
	Dot *ZeroTrustDnsLocationEndpointsDot `field:"optional" json:"dot" yaml:"dot"`
	// ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_dns_location#ipv4 ZeroTrustDnsLocation#ipv4}
	Ipv4 *ZeroTrustDnsLocationEndpointsIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_dns_location#ipv6 ZeroTrustDnsLocation#ipv6}
	Ipv6 *ZeroTrustDnsLocationEndpointsIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

