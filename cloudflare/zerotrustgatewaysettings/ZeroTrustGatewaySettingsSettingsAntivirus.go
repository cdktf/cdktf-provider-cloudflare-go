// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#enabled_download_phase ZeroTrustGatewaySettings#enabled_download_phase}
	EnabledDownloadPhase interface{} `field:"optional" json:"enabledDownloadPhase" yaml:"enabledDownloadPhase"`
	// Enable anti-virus scanning on uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#enabled_upload_phase ZeroTrustGatewaySettings#enabled_upload_phase}
	EnabledUploadPhase interface{} `field:"optional" json:"enabledUploadPhase" yaml:"enabledUploadPhase"`
	// Block requests for files that cannot be scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#fail_closed ZeroTrustGatewaySettings#fail_closed}
	FailClosed interface{} `field:"optional" json:"failClosed" yaml:"failClosed"`
	// Configure a message to display on the user's device when an antivirus search is performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#notification_settings ZeroTrustGatewaySettings#notification_settings}
	NotificationSettings *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
}

