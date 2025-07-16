// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrule


type FirewallRuleActionResponse struct {
	// The response body to return. The value must conform to the configured content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#body FirewallRule#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// The content type of the body. Must be one of the following: `text/plain`, `text/xml`, or `application/json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/firewall_rule#content_type FirewallRule#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

