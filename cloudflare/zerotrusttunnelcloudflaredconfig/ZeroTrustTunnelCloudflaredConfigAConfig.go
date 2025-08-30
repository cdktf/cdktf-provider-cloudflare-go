// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelCloudflaredConfigAConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_tunnel_cloudflared_config#account_id ZeroTrustTunnelCloudflaredConfigA#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// UUID of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_tunnel_cloudflared_config#tunnel_id ZeroTrustTunnelCloudflaredConfigA#tunnel_id}
	TunnelId *string `field:"required" json:"tunnelId" yaml:"tunnelId"`
	// The tunnel configuration and ingress rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_tunnel_cloudflared_config#config ZeroTrustTunnelCloudflaredConfigA#config}
	Config *ZeroTrustTunnelCloudflaredConfigConfig `field:"optional" json:"config" yaml:"config"`
	// Indicates if this is a locally or remotely configured tunnel.
	//
	// If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel's configuration on the Zero Trust dashboard.
	// Available values: "local", "cloudflare".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_tunnel_cloudflared_config#source ZeroTrustTunnelCloudflaredConfigA#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

