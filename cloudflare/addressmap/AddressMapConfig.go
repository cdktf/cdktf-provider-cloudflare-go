// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package addressmap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AddressMapConfig struct {
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
	// Identifier of a Cloudflare account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#account_id AddressMap#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// If you have legacy TLS clients which do not send the TLS server name indicator, then you can specify one default SNI on the map.
	//
	// If Cloudflare receives a TLS handshake from a client without an SNI, it will respond with the default SNI on those IPs. The default SNI can be any valid zone or subdomain owned by the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#default_sni AddressMap#default_sni}
	DefaultSni *string `field:"optional" json:"defaultSni" yaml:"defaultSni"`
	// An optional description field which may be used to describe the types of IPs or zones on the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#description AddressMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the Address Map is enabled or not.
	//
	// Cloudflare's DNS will not respond with IP addresses on an Address Map until the map is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#enabled AddressMap#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#ips AddressMap#ips}.
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map.
	//
	// A zone membership will take priority over an account membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/address_map#memberships AddressMap#memberships}
	Memberships interface{} `field:"optional" json:"memberships" yaml:"memberships"`
}

