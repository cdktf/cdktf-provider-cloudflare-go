// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfigIngressRule struct {
	// Name of the service to which the request will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/resources/zero_trust_tunnel_cloudflared_config#service ZeroTrustTunnelCloudflaredConfigA#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Hostname to match the incoming request with. If the hostname matches, the request will be sent to the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/resources/zero_trust_tunnel_cloudflared_config#hostname ZeroTrustTunnelCloudflaredConfigA#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// origin_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/resources/zero_trust_tunnel_cloudflared_config#origin_request ZeroTrustTunnelCloudflaredConfigA#origin_request}
	OriginRequest *ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// Path of the incoming request. If the path matches, the request will be sent to the local service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.47.0/docs/resources/zero_trust_tunnel_cloudflared_config#path ZeroTrustTunnelCloudflaredConfigA#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

