// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount


type TeamsAccountAntivirusNotificationSettings struct {
	// Enable notification settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.32.0/docs/resources/teams_account#enabled TeamsAccount#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Notification content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.32.0/docs/resources/teams_account#message TeamsAccount#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Support URL to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.32.0/docs/resources/teams_account#support_url TeamsAccount#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

