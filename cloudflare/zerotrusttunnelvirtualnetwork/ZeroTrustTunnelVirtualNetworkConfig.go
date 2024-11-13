// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelvirtualnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustTunnelVirtualNetworkConfig struct {
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
	// The account identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_tunnel_virtual_network#account_id ZeroTrustTunnelVirtualNetwork#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A user-friendly name chosen when the virtual network is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_tunnel_virtual_network#name ZeroTrustTunnelVirtualNetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the tunnel virtual network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_tunnel_virtual_network#comment ZeroTrustTunnelVirtualNetwork#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_tunnel_virtual_network#id ZeroTrustTunnelVirtualNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this virtual network is the default one for the account.
	//
	// This means IP Routes belong to this virtual network and Teams Clients in the account route through this virtual network, unless specified otherwise for each case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/resources/zero_trust_tunnel_virtual_network#is_default_network ZeroTrustTunnelVirtualNetwork#is_default_network}
	IsDefaultNetwork interface{} `field:"optional" json:"isDefaultNetwork" yaml:"isDefaultNetwork"`
}

