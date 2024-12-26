// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdextest


type ZeroTrustDexTestData struct {
	// The host URL for `http` test `kind`. For `traceroute`, it must be a valid hostname or IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_dex_test#host ZeroTrustDexTest#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The type of Device Dex Test. Available values: `http`, `traceroute`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_dex_test#kind ZeroTrustDexTest#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The http request method. Available values: `GET`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_dex_test#method ZeroTrustDexTest#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

