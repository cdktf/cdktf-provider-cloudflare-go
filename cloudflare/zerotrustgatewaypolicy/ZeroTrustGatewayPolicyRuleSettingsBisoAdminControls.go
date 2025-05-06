// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls struct {
	// Configure whether copy is enabled or not.
	//
	// When set with "remote_only", copying isolated content from the remote browser to the user's local clipboard is disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled", "remote_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#copy ZeroTrustGatewayPolicy#copy}
	Copy *string `field:"optional" json:"copy" yaml:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#dcp ZeroTrustGatewayPolicy#dcp}
	Dcp interface{} `field:"optional" json:"dcp" yaml:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#dd ZeroTrustGatewayPolicy#dd}
	Dd interface{} `field:"optional" json:"dd" yaml:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#dk ZeroTrustGatewayPolicy#dk}
	Dk interface{} `field:"optional" json:"dk" yaml:"dk"`
	// Configure whether downloading enabled or not.
	//
	// When absent, downloading is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#download ZeroTrustGatewayPolicy#download}
	Download *string `field:"optional" json:"download" yaml:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#dp ZeroTrustGatewayPolicy#dp}
	Dp interface{} `field:"optional" json:"dp" yaml:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#du ZeroTrustGatewayPolicy#du}
	Du interface{} `field:"optional" json:"du" yaml:"du"`
	// Configure whether keyboard usage is enabled or not.
	//
	// When absent, keyboard usage is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#keyboard ZeroTrustGatewayPolicy#keyboard}
	Keyboard *string `field:"optional" json:"keyboard" yaml:"keyboard"`
	// Configure whether pasting is enabled or not.
	//
	// When set with "remote_only", pasting content from the user's local clipboard into isolated pages is disabled. When absent, paste is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled", "remote_only".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#paste ZeroTrustGatewayPolicy#paste}
	Paste *string `field:"optional" json:"paste" yaml:"paste"`
	// Configure whether printing is enabled or not.
	//
	// When absent, printing is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#printing ZeroTrustGatewayPolicy#printing}
	Printing *string `field:"optional" json:"printing" yaml:"printing"`
	// Configure whether uploading is enabled or not.
	//
	// When absent, uploading is enabled. Only applies when `version == "v2"`.
	// Available values: "enabled", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#upload ZeroTrustGatewayPolicy#upload}
	Upload *string `field:"optional" json:"upload" yaml:"upload"`
	// Indicates which version of the browser isolation controls should apply. Available values: "v1", "v2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#version ZeroTrustGatewayPolicy#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

