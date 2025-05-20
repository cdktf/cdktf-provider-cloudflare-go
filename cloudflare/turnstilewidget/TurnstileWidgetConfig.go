// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package turnstilewidget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TurnstileWidgetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#account_id TurnstileWidget#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#domains TurnstileWidget#domains}.
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// Widget Mode Available values: "non-interactive", "invisible", "managed".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#mode TurnstileWidget#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Human readable widget name.
	//
	// Not unique. Cloudflare suggests that you
	// set this to a meaningful string to make it easier to identify your
	// widget, and where it is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#name TurnstileWidget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive challenges in response to malicious bots (ENT only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#bot_fight_mode TurnstileWidget#bot_fight_mode}
	BotFightMode interface{} `field:"optional" json:"botFightMode" yaml:"botFightMode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance, this setting can determine the clearance level to be set Available values: "no_clearance", "jschallenge", "managed", "interactive".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#clearance_level TurnstileWidget#clearance_level}
	ClearanceLevel *string `field:"optional" json:"clearanceLevel" yaml:"clearanceLevel"`
	// Return the Ephemeral ID in /siteverify (ENT only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#ephemeral_id TurnstileWidget#ephemeral_id}
	EphemeralId interface{} `field:"optional" json:"ephemeralId" yaml:"ephemeralId"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#offlabel TurnstileWidget#offlabel}
	Offlabel interface{} `field:"optional" json:"offlabel" yaml:"offlabel"`
	// Region where this widget can be used. This cannot be changed after creation. Available values: "world", "china".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/turnstile_widget#region TurnstileWidget#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

