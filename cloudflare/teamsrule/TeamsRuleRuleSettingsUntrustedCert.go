// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsUntrustedCert struct {
	// Action to be taken when the SSL certificate of upstream is invalid. Available values: `pass_through`, `block`, `error`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/teams_rule#action TeamsRule#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
}

