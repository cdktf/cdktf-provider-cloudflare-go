// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelwarpconnector


type DataCloudflareZeroTrustTunnelWarpConnectorFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#exclude_prefix DataCloudflareZeroTrustTunnelWarpConnector#exclude_prefix}.
	ExcludePrefix *string `field:"optional" json:"excludePrefix" yaml:"excludePrefix"`
	// If provided, include only resources that were created (and not deleted) before this time. URL encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#existed_at DataCloudflareZeroTrustTunnelWarpConnector#existed_at}
	ExistedAt *string `field:"optional" json:"existedAt" yaml:"existedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#include_prefix DataCloudflareZeroTrustTunnelWarpConnector#include_prefix}.
	IncludePrefix *string `field:"optional" json:"includePrefix" yaml:"includePrefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#is_deleted DataCloudflareZeroTrustTunnelWarpConnector#is_deleted}
	IsDeleted interface{} `field:"optional" json:"isDeleted" yaml:"isDeleted"`
	// A user-friendly name for the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#name DataCloudflareZeroTrustTunnelWarpConnector#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the tunnel.
	//
	// Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	// Available values: "inactive", "degraded", "healthy", "down".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#status DataCloudflareZeroTrustTunnelWarpConnector#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// UUID of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#uuid DataCloudflareZeroTrustTunnelWarpConnector#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#was_active_at DataCloudflareZeroTrustTunnelWarpConnector#was_active_at}.
	WasActiveAt *string `field:"optional" json:"wasActiveAt" yaml:"wasActiveAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/data-sources/zero_trust_tunnel_warp_connector#was_inactive_at DataCloudflareZeroTrustTunnelWarpConnector#was_inactive_at}.
	WasInactiveAt *string `field:"optional" json:"wasInactiveAt" yaml:"wasInactiveAt"`
}

