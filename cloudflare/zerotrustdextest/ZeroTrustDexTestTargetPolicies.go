// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdextest


type ZeroTrustDexTestTargetPolicies struct {
	// Whether the DEX rule is the account default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_dex_test#default ZeroTrustDexTest#default}
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// The id of the DEX rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_dex_test#id ZeroTrustDexTest#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the DEX rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_dex_test#name ZeroTrustDexTest#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

