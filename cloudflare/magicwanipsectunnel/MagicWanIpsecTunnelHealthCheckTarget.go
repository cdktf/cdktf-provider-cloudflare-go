// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanipsectunnel


type MagicWanIpsecTunnelHealthCheckTarget struct {
	// The saved health check target.
	//
	// Setting the value to the empty string indicates that the calculated default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_wan_ipsec_tunnel#saved MagicWanIpsecTunnel#saved}
	Saved *string `field:"optional" json:"saved" yaml:"saved"`
}

