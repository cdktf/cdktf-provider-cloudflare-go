// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule


type EmailRoutingRuleAction struct {
	// Type of action. Available values: `forward`, `worker`, `drop`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.25.0/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value to match on. Required for `type` of `literal`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.25.0/docs/resources/email_routing_rule#value EmailRoutingRule#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

