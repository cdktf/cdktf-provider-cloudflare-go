// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwangretunnel


type MagicWanGreTunnelHealthCheckTarget struct {
	// The saved health check target.
	//
	// Setting the value to the empty string indicates that the calculated default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_wan_gre_tunnel#saved MagicWanGreTunnel#saved}
	Saved *string `field:"optional" json:"saved" yaml:"saved"`
}

