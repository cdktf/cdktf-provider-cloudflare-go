// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule


type EmailRoutingRuleMatcher struct {
	// Type of matcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.15.0/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Field for type matcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.15.0/docs/resources/email_routing_rule#field EmailRoutingRule#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Value for matcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.15.0/docs/resources/email_routing_rule#value EmailRoutingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

