// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfig struct {
	// ingress_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_tunnel_cloudflared_config#ingress_rule ZeroTrustTunnelCloudflaredConfigA#ingress_rule}
	IngressRule interface{} `field:"required" json:"ingressRule" yaml:"ingressRule"`
	// origin_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_tunnel_cloudflared_config#origin_request ZeroTrustTunnelCloudflaredConfigA#origin_request}
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// warp_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_tunnel_cloudflared_config#warp_routing ZeroTrustTunnelCloudflaredConfigA#warp_routing}
	WarpRouting *ZeroTrustTunnelCloudflaredConfigConfigWarpRouting `field:"optional" json:"warpRouting" yaml:"warpRouting"`
}

