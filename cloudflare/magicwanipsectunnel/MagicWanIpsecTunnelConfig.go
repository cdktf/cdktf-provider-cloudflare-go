// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicwanipsectunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicWanIpsecTunnelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#account_id MagicWanIpsecTunnel#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#cloudflare_endpoint MagicWanIpsecTunnel#cloudflare_endpoint}
	CloudflareEndpoint *string `field:"required" json:"cloudflareEndpoint" yaml:"cloudflareEndpoint"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side of the tunnel.
	//
	// Select the subnet from the following private IP space: 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#interface_address MagicWanIpsecTunnel#interface_address}
	InterfaceAddress *string `field:"required" json:"interfaceAddress" yaml:"interfaceAddress"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#name MagicWanIpsecTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IP address assigned to the customer side of the IPsec tunnel.
	//
	// Not required, but must be set for proactive traceroutes to work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#customer_endpoint MagicWanIpsecTunnel#customer_endpoint}
	CustomerEndpoint *string `field:"optional" json:"customerEndpoint" yaml:"customerEndpoint"`
	// An optional description forthe IPsec tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#description MagicWanIpsecTunnel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#health_check MagicWanIpsecTunnel#health_check}.
	HealthCheck *MagicWanIpsecTunnelHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the address being the first IP of the subnet and not same as the address of virtual_subnet6.
	//
	// Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 , interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#interface_address6 MagicWanIpsecTunnel#interface_address6}
	InterfaceAddress6 *string `field:"optional" json:"interfaceAddress6" yaml:"interfaceAddress6"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#psk MagicWanIpsecTunnel#psk}
	Psk *string `field:"optional" json:"psk" yaml:"psk"`
	// If `true`, then IPsec replay protection will be supported in the Cloudflare-to-customer direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/magic_wan_ipsec_tunnel#replay_protection MagicWanIpsecTunnel#replay_protection}
	ReplayProtection interface{} `field:"optional" json:"replayProtection" yaml:"replayProtection"`
}

