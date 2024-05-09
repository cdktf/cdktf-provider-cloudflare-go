// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsDnsResolvers struct {
	// ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.32.0/docs/resources/teams_rule#ipv4 TeamsRule#ipv4}
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.32.0/docs/resources/teams_rule#ipv6 TeamsRule#ipv6}
	Ipv6 interface{} `field:"optional" json:"ipv6" yaml:"ipv6"`
}

