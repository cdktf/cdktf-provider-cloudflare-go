// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyConnectionRules struct {
	// ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_access_policy#ssh ZeroTrustAccessPolicy#ssh}
	Ssh *ZeroTrustAccessPolicyConnectionRulesSsh `field:"required" json:"ssh" yaml:"ssh"`
}

