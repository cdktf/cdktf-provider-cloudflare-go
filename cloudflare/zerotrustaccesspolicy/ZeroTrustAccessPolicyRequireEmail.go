// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyRequireEmail struct {
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

