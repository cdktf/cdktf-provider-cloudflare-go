// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfigIngressRuleOriginRequestAccess struct {
	// Audience tags of the access rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#aud_tag ZeroTrustTunnelCloudflaredConfigA#aud_tag}
	AudTag *[]*string `field:"optional" json:"audTag" yaml:"audTag"`
	// Whether the access rule is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#required ZeroTrustTunnelCloudflaredConfigA#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Name of the team to which the access rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_tunnel_cloudflared_config#team_name ZeroTrustTunnelCloudflaredConfigA#team_name}
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
}

