// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS. 140-2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_gateway_settings#tls ZeroTrustGatewaySettings#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

