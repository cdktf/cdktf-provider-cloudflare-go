// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettings struct {
	// Activity log settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#activity_log ZeroTrustGatewaySettings#activity_log}
	ActivityLog *ZeroTrustGatewaySettingsSettingsActivityLog `field:"optional" json:"activityLog" yaml:"activityLog"`
	// Anti-virus settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#antivirus ZeroTrustGatewaySettings#antivirus}
	Antivirus *ZeroTrustGatewaySettingsSettingsAntivirus `field:"optional" json:"antivirus" yaml:"antivirus"`
	// Block page layout settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#block_page ZeroTrustGatewaySettings#block_page}
	BlockPage *ZeroTrustGatewaySettingsSettingsBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// DLP body scanning settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#body_scanning ZeroTrustGatewaySettings#body_scanning}
	BodyScanning *ZeroTrustGatewaySettingsSettingsBodyScanning `field:"optional" json:"bodyScanning" yaml:"bodyScanning"`
	// Browser isolation settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#browser_isolation ZeroTrustGatewaySettings#browser_isolation}
	BrowserIsolation *ZeroTrustGatewaySettingsSettingsBrowserIsolation `field:"optional" json:"browserIsolation" yaml:"browserIsolation"`
	// Certificate settings for Gateway TLS interception. If not specified, the Cloudflare Root CA will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#certificate ZeroTrustGatewaySettings#certificate}
	Certificate *ZeroTrustGatewaySettingsSettingsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by `certificate`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#custom_certificate ZeroTrustGatewaySettings#custom_certificate}
	CustomCertificate *ZeroTrustGatewaySettingsSettingsCustomCertificate `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Extended e-mail matching settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#extended_email_matching ZeroTrustGatewaySettings#extended_email_matching}
	ExtendedEmailMatching *ZeroTrustGatewaySettingsSettingsExtendedEmailMatching `field:"optional" json:"extendedEmailMatching" yaml:"extendedEmailMatching"`
	// FIPS settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#fips ZeroTrustGatewaySettings#fips}
	Fips *ZeroTrustGatewaySettingsSettingsFips `field:"optional" json:"fips" yaml:"fips"`
	// Setting to enable host selector in egress policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#host_selector ZeroTrustGatewaySettings#host_selector}
	HostSelector *ZeroTrustGatewaySettingsSettingsHostSelector `field:"optional" json:"hostSelector" yaml:"hostSelector"`
	// Protocol Detection settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#protocol_detection ZeroTrustGatewaySettings#protocol_detection}
	ProtocolDetection *ZeroTrustGatewaySettingsSettingsProtocolDetection `field:"optional" json:"protocolDetection" yaml:"protocolDetection"`
	// Sandbox settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#sandbox ZeroTrustGatewaySettings#sandbox}
	Sandbox *ZeroTrustGatewaySettingsSettingsSandbox `field:"optional" json:"sandbox" yaml:"sandbox"`
	// TLS interception settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_settings#tls_decrypt ZeroTrustGatewaySettings#tls_decrypt}
	TlsDecrypt *ZeroTrustGatewaySettingsSettingsTlsDecrypt `field:"optional" json:"tlsDecrypt" yaml:"tlsDecrypt"`
}

