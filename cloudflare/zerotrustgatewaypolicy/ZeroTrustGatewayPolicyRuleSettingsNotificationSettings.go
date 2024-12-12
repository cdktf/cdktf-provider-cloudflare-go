// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsNotificationSettings struct {
	// Enable notification settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_gateway_policy#enabled ZeroTrustGatewayPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Notification content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_gateway_policy#message ZeroTrustGatewayPolicy#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Support URL to show in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.48.0/docs/resources/zero_trust_gateway_policy#support_url ZeroTrustGatewayPolicy#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

