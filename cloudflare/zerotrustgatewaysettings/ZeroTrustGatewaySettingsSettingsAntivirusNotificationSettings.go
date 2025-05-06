// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings struct {
	// Set notification on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If true, context information will be passed as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#include_context ZeroTrustGatewaySettings#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// Customize the message shown in the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#msg ZeroTrustGatewaySettings#msg}
	Msg *string `field:"optional" json:"msg" yaml:"msg"`
	// Optional URL to direct users to additional information. If not set, the notification will open a block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#support_url ZeroTrustGatewaySettings#support_url}
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
}

