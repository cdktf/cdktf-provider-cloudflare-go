// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationEndpointsIpv4 struct {
	// True if the endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dns_location#enabled ZeroTrustDnsLocation#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

