// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesConnectionRulesSsh struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#usernames ZeroTrustAccessApplication#usernames}
	Usernames *[]*string `field:"required" json:"usernames" yaml:"usernames"`
	// Enables using Identity Provider email alias as SSH username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#allow_email_alias ZeroTrustAccessApplication#allow_email_alias}
	AllowEmailAlias interface{} `field:"optional" json:"allowEmailAlias" yaml:"allowEmailAlias"`
}

