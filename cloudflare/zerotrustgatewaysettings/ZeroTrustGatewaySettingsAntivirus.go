// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsAntivirus struct {
	// Scan on file download.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_gateway_settings#enabled_download_phase ZeroTrustGatewaySettings#enabled_download_phase}
	EnabledDownloadPhase interface{} `field:"required" json:"enabledDownloadPhase" yaml:"enabledDownloadPhase"`
	// Scan on file upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_gateway_settings#enabled_upload_phase ZeroTrustGatewaySettings#enabled_upload_phase}
	EnabledUploadPhase interface{} `field:"required" json:"enabledUploadPhase" yaml:"enabledUploadPhase"`
	// Block requests for files that cannot be scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_gateway_settings#fail_closed ZeroTrustGatewaySettings#fail_closed}
	FailClosed interface{} `field:"required" json:"failClosed" yaml:"failClosed"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_gateway_settings#notification_settings ZeroTrustGatewaySettings#notification_settings}
	NotificationSettings *ZeroTrustGatewaySettingsAntivirusNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
}

