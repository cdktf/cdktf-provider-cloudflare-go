// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroomrules


type WaitingRoomRulesRules struct {
	// The action to take when the expression matches. Available values: "bypass_waiting_room".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/waiting_room_rules#action WaitingRoomRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Criteria defining when there is a match for the current rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/waiting_room_rules#expression WaitingRoomRules#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/waiting_room_rules#description WaitingRoomRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// When set to true, the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/waiting_room_rules#enabled WaitingRoomRules#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

