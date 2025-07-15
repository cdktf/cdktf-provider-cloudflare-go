// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsIpv6 struct {
	// True if the endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_dns_location#enabled ZeroTrustDnsLocation#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint.
	//
	// When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

