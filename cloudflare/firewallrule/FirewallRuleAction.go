// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrule


type FirewallRuleAction struct {
	// The action to perform. Available values: "simulate", "ban", "challenge", "js_challenge", "managed_challenge".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/firewall_rule#mode FirewallRule#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded.
	//
	// The custom response configured in this object will override the custom error for the zone. This object is optional.
	// Notes: If you omit this object, Cloudflare will use the default HTML error page. If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone challenge pages and you should not provide the "response" object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/firewall_rule#response FirewallRule#response}
	Response *FirewallRuleActionResponse `field:"optional" json:"response" yaml:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	//
	// Must be an integer value greater than or equal to the period.
	// Notes: If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/firewall_rule#timeout FirewallRule#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

