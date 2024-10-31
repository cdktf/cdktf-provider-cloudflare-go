// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyConnectionRulesSsh struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/zero_trust_access_policy#usernames ZeroTrustAccessPolicy#usernames}
	Usernames *[]*string `field:"required" json:"usernames" yaml:"usernames"`
}

