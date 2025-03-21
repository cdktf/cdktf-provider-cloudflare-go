// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package web3hostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Web3HostnameConfig struct {
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
	// The hostname that will point to the target gateway via CNAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/web3_hostname#name Web3Hostname#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target gateway of the hostname. Available values: "ethereum", "ipfs", "ipfs_universal_path".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/web3_hostname#target Web3Hostname#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/web3_hostname#zone_id Web3Hostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// An optional description of the hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/web3_hostname#description Web3Hostname#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// DNSLink value used if the target is ipfs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/web3_hostname#dnslink Web3Hostname#dnslink}
	Dnslink *string `field:"optional" json:"dnslink" yaml:"dnslink"`
}

