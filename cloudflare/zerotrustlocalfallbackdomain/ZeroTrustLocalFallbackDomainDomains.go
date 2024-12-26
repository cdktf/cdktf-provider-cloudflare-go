// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustlocalfallbackdomain


type ZeroTrustLocalFallbackDomainDomains struct {
	// A description of the fallback domain, displayed in the client UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_local_fallback_domain#description ZeroTrustLocalFallbackDomain#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of IP addresses to handle domain resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_local_fallback_domain#dns_server ZeroTrustLocalFallbackDomain#dns_server}
	DnsServer *[]*string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
	// The domain suffix to match when resolving locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_local_fallback_domain#suffix ZeroTrustLocalFallbackDomain#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

