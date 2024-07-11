// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsNotificationSettings struct {
	// Enable notification settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/teams_rule#enabled TeamsRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Notification content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/teams_rule#message TeamsRule#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Support URL to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/teams_rule#support_url TeamsRule#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

