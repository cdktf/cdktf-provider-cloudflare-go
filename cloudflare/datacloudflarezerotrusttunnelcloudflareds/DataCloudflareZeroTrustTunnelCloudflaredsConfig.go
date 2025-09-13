// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrusttunnelcloudflareds

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareZeroTrustTunnelCloudflaredsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#account_id DataCloudflareZeroTrustTunnelCloudflareds#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#exclude_prefix DataCloudflareZeroTrustTunnelCloudflareds#exclude_prefix}.
	ExcludePrefix *string `field:"optional" json:"excludePrefix" yaml:"excludePrefix"`
	// If provided, include only resources that were created (and not deleted) before this time. URL encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#existed_at DataCloudflareZeroTrustTunnelCloudflareds#existed_at}
	ExistedAt *string `field:"optional" json:"existedAt" yaml:"existedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#include_prefix DataCloudflareZeroTrustTunnelCloudflareds#include_prefix}.
	IncludePrefix *string `field:"optional" json:"includePrefix" yaml:"includePrefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#is_deleted DataCloudflareZeroTrustTunnelCloudflareds#is_deleted}
	IsDeleted interface{} `field:"optional" json:"isDeleted" yaml:"isDeleted"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#max_items DataCloudflareZeroTrustTunnelCloudflareds#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// A user-friendly name for a tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#name DataCloudflareZeroTrustTunnelCloudflareds#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The status of the tunnel.
	//
	// Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	// Available values: "inactive", "degraded", "healthy", "down".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#status DataCloudflareZeroTrustTunnelCloudflareds#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// UUID of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#uuid DataCloudflareZeroTrustTunnelCloudflareds#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#was_active_at DataCloudflareZeroTrustTunnelCloudflareds#was_active_at}.
	WasActiveAt *string `field:"optional" json:"wasActiveAt" yaml:"wasActiveAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/data-sources/zero_trust_tunnel_cloudflareds#was_inactive_at DataCloudflareZeroTrustTunnelCloudflareds#was_inactive_at}.
	WasInactiveAt *string `field:"optional" json:"wasInactiveAt" yaml:"wasInactiveAt"`
}

