// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount


type TeamsAccountExtendedEmailMatching struct {
	// Whether e-mails should be matched on all variants of user emails (with + or . modifiers) in Firewall policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#enabled TeamsAccount#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

