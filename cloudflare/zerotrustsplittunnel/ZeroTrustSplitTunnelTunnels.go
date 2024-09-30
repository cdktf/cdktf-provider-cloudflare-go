// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustsplittunnel


type ZeroTrustSplitTunnelTunnels struct {
	// The address for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_split_tunnel#address ZeroTrustSplitTunnel#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// A description for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_split_tunnel#description ZeroTrustSplitTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain name for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/zero_trust_split_tunnel#host ZeroTrustSplitTunnel#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
}

