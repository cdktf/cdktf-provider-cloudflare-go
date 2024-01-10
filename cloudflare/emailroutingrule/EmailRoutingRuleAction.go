// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule


type EmailRoutingRuleAction struct {
	// Type of supported action. Available values: `forward`, `worker`, `drop`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.22.0/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An array with items in the following form. Only required when `type` is `forward` or `worker`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.22.0/docs/resources/email_routing_rule#value EmailRoutingRule#value}
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

