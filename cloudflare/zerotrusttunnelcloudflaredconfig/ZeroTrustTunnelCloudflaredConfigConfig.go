// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be defined for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_tunnel_cloudflared_config#ingress ZeroTrustTunnelCloudflaredConfigA#ingress}
	Ingress interface{} `field:"optional" json:"ingress" yaml:"ingress"`
	// Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_tunnel_cloudflared_config#origin_request ZeroTrustTunnelCloudflaredConfigA#origin_request}
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// Enable private network access from WARP users to private network routes.
	//
	// This is enabled if the tunnel has an assigned route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_tunnel_cloudflared_config#warp_routing ZeroTrustTunnelCloudflaredConfigA#warp_routing}
	WarpRouting *ZeroTrustTunnelCloudflaredConfigConfigWarpRouting `field:"optional" json:"warpRouting" yaml:"warpRouting"`
}

