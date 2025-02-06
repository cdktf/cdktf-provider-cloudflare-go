// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls struct {
	// Configure whether copy is enabled or not.
	//
	// When set with 'remote_only', copying isolated content from the remote browser to the user's local clipboard is disabled. When absent, copy is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#copy ZeroTrustGatewayPolicy#copy}
	Copy *string `field:"optional" json:"copy" yaml:"copy"`
	// Disable clipboard redirection. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_clipboard_redirection ZeroTrustGatewayPolicy#disable_clipboard_redirection}
	DisableClipboardRedirection interface{} `field:"optional" json:"disableClipboardRedirection" yaml:"disableClipboardRedirection"`
	// Disable copy-paste. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_copy_paste ZeroTrustGatewayPolicy#disable_copy_paste}
	DisableCopyPaste interface{} `field:"optional" json:"disableCopyPaste" yaml:"disableCopyPaste"`
	// Disable download. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_download ZeroTrustGatewayPolicy#disable_download}
	DisableDownload interface{} `field:"optional" json:"disableDownload" yaml:"disableDownload"`
	// Disable keyboard usage. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_keyboard ZeroTrustGatewayPolicy#disable_keyboard}
	DisableKeyboard interface{} `field:"optional" json:"disableKeyboard" yaml:"disableKeyboard"`
	// Disable printing. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_printing ZeroTrustGatewayPolicy#disable_printing}
	DisablePrinting interface{} `field:"optional" json:"disablePrinting" yaml:"disablePrinting"`
	// Disable upload. Only applies when version == v1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#disable_upload ZeroTrustGatewayPolicy#disable_upload}
	DisableUpload interface{} `field:"optional" json:"disableUpload" yaml:"disableUpload"`
	// Configure whether downloading enabled or not. When absent, downloading is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#download ZeroTrustGatewayPolicy#download}
	Download *string `field:"optional" json:"download" yaml:"download"`
	// Configure whether keyboard usage is enabled or not.
	//
	// When absent, keyboard usage is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#keyboard ZeroTrustGatewayPolicy#keyboard}
	Keyboard *string `field:"optional" json:"keyboard" yaml:"keyboard"`
	// Configure whether pasting is enabled or not.
	//
	// When set with 'remote_only', pasting content from the user's local clipboard into isolated pages is disabled. When absent, paste is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#paste ZeroTrustGatewayPolicy#paste}
	Paste *string `field:"optional" json:"paste" yaml:"paste"`
	// Configure whether printing is enabled or not. When absent, printing is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#printing ZeroTrustGatewayPolicy#printing}
	Printing *string `field:"optional" json:"printing" yaml:"printing"`
	// Configure whether uploading is enabled or not. When absent, uploading is enabled. Only applies when version == v2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#upload ZeroTrustGatewayPolicy#upload}
	Upload *string `field:"optional" json:"upload" yaml:"upload"`
	// Indicates which version (v1 or v2) of the browser isolation controls should apply. Defaults to `v1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_gateway_policy#version ZeroTrustGatewayPolicy#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

