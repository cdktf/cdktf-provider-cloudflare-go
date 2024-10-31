// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation


type ZeroTrustDnsLocationNetworks struct {
	// CIDR notation representation of the network IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_dns_location#network ZeroTrustDnsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

