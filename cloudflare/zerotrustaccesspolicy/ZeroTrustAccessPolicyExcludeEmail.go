// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeEmail struct {
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

