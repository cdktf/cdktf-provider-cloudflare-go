// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally struct {
	// The fallback behavior to apply when the internal DNS response code is different from 'NOERROR' or when the response data only contains CNAME records for 'A' or 'AAAA' queries.
	//
	// Available values: "none", "public_dns".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#fallback ZeroTrustGatewayPolicy#fallback}
	Fallback *string `field:"optional" json:"fallback" yaml:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_gateway_policy#view_id ZeroTrustGatewayPolicy#view_id}
	ViewId *string `field:"optional" json:"viewId" yaml:"viewId"`
}

