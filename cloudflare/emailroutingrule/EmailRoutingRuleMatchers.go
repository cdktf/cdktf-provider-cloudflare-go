// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailroutingrule


type EmailRoutingRuleMatchers struct {
	// Field for type matcher. Available values: "to".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/email_routing_rule#field EmailRoutingRule#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Type of matcher. Available values: "literal".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/email_routing_rule#type EmailRoutingRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Value for matcher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/email_routing_rule#value EmailRoutingRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

