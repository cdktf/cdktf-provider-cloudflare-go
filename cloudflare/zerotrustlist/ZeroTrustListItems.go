// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustlist


type ZeroTrustListItems struct {
	// The description of the list item, if present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_list#description ZeroTrustList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the item in a list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_list#value ZeroTrustList#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

