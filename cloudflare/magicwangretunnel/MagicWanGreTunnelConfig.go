// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwangretunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicWanGreTunnelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#account_id MagicWanGreTunnel#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#cloudflare_gre_endpoint MagicWanGreTunnel#cloudflare_gre_endpoint}
	CloudflareGreEndpoint *string `field:"optional" json:"cloudflareGreEndpoint" yaml:"cloudflareGreEndpoint"`
	// The IP address assigned to the customer side of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#customer_gre_endpoint MagicWanGreTunnel#customer_gre_endpoint}
	CustomerGreEndpoint *string `field:"optional" json:"customerGreEndpoint" yaml:"customerGreEndpoint"`
	// An optional description of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#description MagicWanGreTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#gre_tunnel_id MagicWanGreTunnel#gre_tunnel_id}
	GreTunnelId *string `field:"optional" json:"greTunnelId" yaml:"greTunnelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#health_check MagicWanGreTunnel#health_check}.
	HealthCheck *MagicWanGreTunnelHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel.
	//
	// Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#interface_address MagicWanGreTunnel#interface_address}
	InterfaceAddress *string `field:"optional" json:"interfaceAddress" yaml:"interfaceAddress"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value is 576.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#mtu MagicWanGreTunnel#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// The name of the tunnel.
	//
	// The name cannot contain spaces or special characters, must be 15 characters or less, and cannot share a name with another GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#name MagicWanGreTunnel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/magic_wan_gre_tunnel#ttl MagicWanGreTunnel#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

