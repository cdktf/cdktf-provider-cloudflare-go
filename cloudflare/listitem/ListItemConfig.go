// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package listitem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ListItemConfig struct {
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
	// The Account ID for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#account_id ListItem#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The unique ID of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#list_id ListItem#list_id}
	ListId *string `field:"required" json:"listId" yaml:"listId"`
	// A non-negative 32 bit integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#asn ListItem#asn}
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// An informative summary of the list item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#comment ListItem#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#hostname ListItem#hostname}
	Hostname *ListItemHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#ip ListItem#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The definition of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/list_item#redirect ListItem#redirect}
	Redirect *ListItemRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

