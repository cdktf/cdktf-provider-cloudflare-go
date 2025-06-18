// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeServiceToken struct {
	// The ID of a Service Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/zero_trust_access_policy#token_id ZeroTrustAccessPolicy#token_id}
	TokenId *string `field:"required" json:"tokenId" yaml:"tokenId"`
}

