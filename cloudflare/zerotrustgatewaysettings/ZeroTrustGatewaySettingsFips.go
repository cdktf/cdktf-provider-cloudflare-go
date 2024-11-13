// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsFips struct {
	// Only allow FIPS-compliant TLS configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_gateway_settings#tls ZeroTrustGatewaySettings#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

