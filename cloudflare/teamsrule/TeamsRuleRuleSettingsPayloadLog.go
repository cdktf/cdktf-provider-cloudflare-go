// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsPayloadLog struct {
	// Enable or disable DLP Payload Logging for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.18.0/docs/resources/teams_rule#enabled TeamsRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

