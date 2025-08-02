// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dns_location#doh ZeroTrustDnsLocation#doh}.
	Doh *ZeroTrustDnsLocationEndpointsDoh `field:"required" json:"doh" yaml:"doh"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dns_location#dot ZeroTrustDnsLocation#dot}.
	Dot *ZeroTrustDnsLocationEndpointsDot `field:"required" json:"dot" yaml:"dot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dns_location#ipv4 ZeroTrustDnsLocation#ipv4}.
	Ipv4 *ZeroTrustDnsLocationEndpointsIpv4 `field:"required" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_dns_location#ipv6 ZeroTrustDnsLocation#ipv6}.
	Ipv6 *ZeroTrustDnsLocationEndpointsIpv6 `field:"required" json:"ipv6" yaml:"ipv6"`
}

