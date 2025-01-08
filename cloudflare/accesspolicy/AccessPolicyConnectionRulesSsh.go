// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyConnectionRulesSsh struct {
	// Contains the Unix usernames that may be used when connecting over SSH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_policy#usernames AccessPolicy#usernames}
	Usernames *[]*string `field:"required" json:"usernames" yaml:"usernames"`
	// Allows connecting to Unix username that matches the authenticating email prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_policy#allow_email_alias AccessPolicy#allow_email_alias}
	AllowEmailAlias interface{} `field:"optional" json:"allowEmailAlias" yaml:"allowEmailAlias"`
}

