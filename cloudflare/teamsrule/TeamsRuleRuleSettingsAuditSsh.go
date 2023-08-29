// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsAuditSsh struct {
	// Log all SSH commands.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.13.0/docs/resources/teams_rule#command_logging TeamsRule#command_logging}
	CommandLogging interface{} `field:"required" json:"commandLogging" yaml:"commandLogging"`
}

