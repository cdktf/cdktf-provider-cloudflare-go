// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflaredvirtualnetwork


type DataCloudflareZeroTrustTunnelCloudflaredVirtualNetworkFilter struct {
	// UUID of the virtual network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_cloudflared_virtual_network#id DataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If `true`, only include the default virtual network.
	//
	// If `false`, exclude the default virtual network. If empty, all virtual networks will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_cloudflared_virtual_network#is_default DataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// If `true`, only include the default virtual network.
	//
	// If `false`, exclude the default virtual network. If empty, all virtual networks will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_cloudflared_virtual_network#is_default_network DataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork#is_default_network}
	IsDefaultNetwork interface{} `field:"optional" json:"isDefaultNetwork" yaml:"isDefaultNetwork"`
	// If `true`, only include deleted virtual networks.
	//
	// If `false`, exclude deleted virtual networks. If empty, all virtual networks will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_cloudflared_virtual_network#is_deleted DataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork#is_deleted}
	IsDeleted interface{} `field:"optional" json:"isDeleted" yaml:"isDeleted"`
	// A user-friendly name for the virtual network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_cloudflared_virtual_network#name DataCloudflareZeroTrustTunnelCloudflaredVirtualNetwork#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

