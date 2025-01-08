// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsPayloadLog struct {
	// Public key used to encrypt matched payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_gateway_settings#public_key ZeroTrustGatewaySettings#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

