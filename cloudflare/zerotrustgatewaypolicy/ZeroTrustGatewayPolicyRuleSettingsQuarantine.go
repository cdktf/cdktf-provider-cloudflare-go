// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsQuarantine struct {
	// Types of files to sandbox.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_gateway_policy#file_types ZeroTrustGatewayPolicy#file_types}
	FileTypes *[]*string `field:"optional" json:"fileTypes" yaml:"fileTypes"`
}

