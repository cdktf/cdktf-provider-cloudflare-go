// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesConnectionRules struct {
	// The SSH-specific rules that define how users may connect to the targets secured by your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_access_application#ssh ZeroTrustAccessApplication#ssh}
	Ssh *ZeroTrustAccessApplicationPoliciesConnectionRulesSsh `field:"optional" json:"ssh" yaml:"ssh"`
}

