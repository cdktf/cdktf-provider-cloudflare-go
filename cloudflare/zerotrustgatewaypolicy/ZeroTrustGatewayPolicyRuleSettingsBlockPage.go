// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsBlockPage struct {
	// URI to which the user will be redirected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_gateway_policy#target_uri ZeroTrustGatewayPolicy#target_uri}
	TargetUri *string `field:"required" json:"targetUri" yaml:"targetUri"`
	// If true, context information will be passed as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_gateway_policy#include_context ZeroTrustGatewayPolicy#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
}

