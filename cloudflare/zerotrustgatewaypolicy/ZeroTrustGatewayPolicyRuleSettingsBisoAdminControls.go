// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls struct {
	// Disable clipboard redirection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_clipboard_redirection ZeroTrustGatewayPolicy#disable_clipboard_redirection}
	DisableClipboardRedirection interface{} `field:"optional" json:"disableClipboardRedirection" yaml:"disableClipboardRedirection"`
	// Disable copy-paste.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_copy_paste ZeroTrustGatewayPolicy#disable_copy_paste}
	DisableCopyPaste interface{} `field:"optional" json:"disableCopyPaste" yaml:"disableCopyPaste"`
	// Disable download.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_download ZeroTrustGatewayPolicy#disable_download}
	DisableDownload interface{} `field:"optional" json:"disableDownload" yaml:"disableDownload"`
	// Disable keyboard usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_keyboard ZeroTrustGatewayPolicy#disable_keyboard}
	DisableKeyboard interface{} `field:"optional" json:"disableKeyboard" yaml:"disableKeyboard"`
	// Disable printing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_printing ZeroTrustGatewayPolicy#disable_printing}
	DisablePrinting interface{} `field:"optional" json:"disablePrinting" yaml:"disablePrinting"`
	// Disable upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/zero_trust_gateway_policy#disable_upload ZeroTrustGatewayPolicy#disable_upload}
	DisableUpload interface{} `field:"optional" json:"disableUpload" yaml:"disableUpload"`
}

