// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamsAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#account_id TeamsAccount#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether to enable the activity log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#activity_log_enabled TeamsAccount#activity_log_enabled}
	ActivityLogEnabled interface{} `field:"optional" json:"activityLogEnabled" yaml:"activityLogEnabled"`
	// antivirus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#antivirus TeamsAccount#antivirus}
	Antivirus *TeamsAccountAntivirus `field:"optional" json:"antivirus" yaml:"antivirus"`
	// block_page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#block_page TeamsAccount#block_page}
	BlockPage *TeamsAccountBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// body_scanning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#body_scanning TeamsAccount#body_scanning}
	BodyScanning *TeamsAccountBodyScanning `field:"optional" json:"bodyScanning" yaml:"bodyScanning"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#certificate TeamsAccount#certificate}
	Certificate *TeamsAccountCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// custom_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#custom_certificate TeamsAccount#custom_certificate}
	CustomCertificate *TeamsAccountCustomCertificate `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// extended_email_matching block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#extended_email_matching TeamsAccount#extended_email_matching}
	ExtendedEmailMatching *TeamsAccountExtendedEmailMatching `field:"optional" json:"extendedEmailMatching" yaml:"extendedEmailMatching"`
	// fips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#fips TeamsAccount#fips}
	Fips *TeamsAccountFips `field:"optional" json:"fips" yaml:"fips"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#id TeamsAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#logging TeamsAccount#logging}
	Logging *TeamsAccountLogging `field:"optional" json:"logging" yaml:"logging"`
	// Enable non-identity onramp for Browser Isolation. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#non_identity_browser_isolation_enabled TeamsAccount#non_identity_browser_isolation_enabled}
	NonIdentityBrowserIsolationEnabled interface{} `field:"optional" json:"nonIdentityBrowserIsolationEnabled" yaml:"nonIdentityBrowserIsolationEnabled"`
	// payload_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#payload_log TeamsAccount#payload_log}
	PayloadLog *TeamsAccountPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// Indicator that protocol detection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#protocol_detection_enabled TeamsAccount#protocol_detection_enabled}
	ProtocolDetectionEnabled interface{} `field:"optional" json:"protocolDetectionEnabled" yaml:"protocolDetectionEnabled"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#proxy TeamsAccount#proxy}
	Proxy *TeamsAccountProxy `field:"optional" json:"proxy" yaml:"proxy"`
	// ssh_session_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#ssh_session_log TeamsAccount#ssh_session_log}
	SshSessionLog *TeamsAccountSshSessionLog `field:"optional" json:"sshSessionLog" yaml:"sshSessionLog"`
	// Indicator that decryption of TLS traffic is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#tls_decrypt_enabled TeamsAccount#tls_decrypt_enabled}
	TlsDecryptEnabled interface{} `field:"optional" json:"tlsDecryptEnabled" yaml:"tlsDecryptEnabled"`
	// Safely browse websites in Browser Isolation through a URL. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/teams_account#url_browser_isolation_enabled TeamsAccount#url_browser_isolation_enabled}
	UrlBrowserIsolationEnabled interface{} `field:"optional" json:"urlBrowserIsolationEnabled" yaml:"urlBrowserIsolationEnabled"`
}

