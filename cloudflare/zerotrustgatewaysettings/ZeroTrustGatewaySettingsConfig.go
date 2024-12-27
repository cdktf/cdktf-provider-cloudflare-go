// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewaySettingsConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#account_id ZeroTrustGatewaySettings#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether to enable the activity log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#activity_log_enabled ZeroTrustGatewaySettings#activity_log_enabled}
	ActivityLogEnabled interface{} `field:"optional" json:"activityLogEnabled" yaml:"activityLogEnabled"`
	// antivirus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#antivirus ZeroTrustGatewaySettings#antivirus}
	Antivirus *ZeroTrustGatewaySettingsAntivirus `field:"optional" json:"antivirus" yaml:"antivirus"`
	// block_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#block_page ZeroTrustGatewaySettings#block_page}
	BlockPage *ZeroTrustGatewaySettingsBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// body_scanning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#body_scanning ZeroTrustGatewaySettings#body_scanning}
	BodyScanning *ZeroTrustGatewaySettingsBodyScanning `field:"optional" json:"bodyScanning" yaml:"bodyScanning"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#certificate ZeroTrustGatewaySettings#certificate}
	Certificate *ZeroTrustGatewaySettingsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// custom_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#custom_certificate ZeroTrustGatewaySettings#custom_certificate}
	CustomCertificate *ZeroTrustGatewaySettingsCustomCertificate `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// extended_email_matching block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#extended_email_matching ZeroTrustGatewaySettings#extended_email_matching}
	ExtendedEmailMatching *ZeroTrustGatewaySettingsExtendedEmailMatching `field:"optional" json:"extendedEmailMatching" yaml:"extendedEmailMatching"`
	// fips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#fips ZeroTrustGatewaySettings#fips}
	Fips *ZeroTrustGatewaySettingsFips `field:"optional" json:"fips" yaml:"fips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#id ZeroTrustGatewaySettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#logging ZeroTrustGatewaySettings#logging}
	Logging *ZeroTrustGatewaySettingsLogging `field:"optional" json:"logging" yaml:"logging"`
	// Enable non-identity onramp for Browser Isolation. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#non_identity_browser_isolation_enabled ZeroTrustGatewaySettings#non_identity_browser_isolation_enabled}
	NonIdentityBrowserIsolationEnabled interface{} `field:"optional" json:"nonIdentityBrowserIsolationEnabled" yaml:"nonIdentityBrowserIsolationEnabled"`
	// payload_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#payload_log ZeroTrustGatewaySettings#payload_log}
	PayloadLog *ZeroTrustGatewaySettingsPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// Indicator that protocol detection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#protocol_detection_enabled ZeroTrustGatewaySettings#protocol_detection_enabled}
	ProtocolDetectionEnabled interface{} `field:"optional" json:"protocolDetectionEnabled" yaml:"protocolDetectionEnabled"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#proxy ZeroTrustGatewaySettings#proxy}
	Proxy *ZeroTrustGatewaySettingsProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// ssh_session_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#ssh_session_log ZeroTrustGatewaySettings#ssh_session_log}
	SshSessionLog *ZeroTrustGatewaySettingsSshSessionLog `field:"optional" json:"sshSessionLog" yaml:"sshSessionLog"`
	// Indicator that decryption of TLS traffic is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#tls_decrypt_enabled ZeroTrustGatewaySettings#tls_decrypt_enabled}
	TlsDecryptEnabled interface{} `field:"optional" json:"tlsDecryptEnabled" yaml:"tlsDecryptEnabled"`
	// Safely browse websites in Browser Isolation through a URL. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_settings#url_browser_isolation_enabled ZeroTrustGatewaySettings#url_browser_isolation_enabled}
	UrlBrowserIsolationEnabled interface{} `field:"optional" json:"urlBrowserIsolationEnabled" yaml:"urlBrowserIsolationEnabled"`
}

