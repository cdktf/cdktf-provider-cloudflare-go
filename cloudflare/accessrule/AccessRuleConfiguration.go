// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessrule


type AccessRuleConfiguration struct {
	// The configuration target.
	//
	// You must set the target to `ip` when specifying an IP address in the rule.
	// Available values: "ip", "ip6", "ip_range", "asn", "country".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/access_rule#target AccessRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The IP address to match. This address will be compared to the IP address of incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/access_rule#value AccessRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

