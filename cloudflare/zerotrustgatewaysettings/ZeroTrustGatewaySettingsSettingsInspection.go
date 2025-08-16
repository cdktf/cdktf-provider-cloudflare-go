// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsInspection struct {
	// Defines the mode of inspection the proxy will use.
	//
	// - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
	// - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
	// Available values: "static", "dynamic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_settings#mode ZeroTrustGatewaySettings#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

