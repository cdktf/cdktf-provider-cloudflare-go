// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdextest


type ZeroTrustDexTestData struct {
	// The desired endpoint to test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_dex_test#host ZeroTrustDexTest#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The type of test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_dex_test#kind ZeroTrustDexTest#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// The HTTP request method type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_dex_test#method ZeroTrustDexTest#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

