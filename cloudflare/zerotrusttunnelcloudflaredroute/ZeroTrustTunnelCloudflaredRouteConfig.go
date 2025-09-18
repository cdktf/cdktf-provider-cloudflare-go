// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelCloudflaredRouteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Cloudflare account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_tunnel_cloudflared_route#account_id ZeroTrustTunnelCloudflaredRoute#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_tunnel_cloudflared_route#network ZeroTrustTunnelCloudflaredRoute#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// UUID of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_tunnel_cloudflared_route#tunnel_id ZeroTrustTunnelCloudflaredRoute#tunnel_id}
	TunnelId *string `field:"required" json:"tunnelId" yaml:"tunnelId"`
	// Optional remark describing the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_tunnel_cloudflared_route#comment ZeroTrustTunnelCloudflaredRoute#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// UUID of the virtual network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_tunnel_cloudflared_route#virtual_network_id ZeroTrustTunnelCloudflaredRoute#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

