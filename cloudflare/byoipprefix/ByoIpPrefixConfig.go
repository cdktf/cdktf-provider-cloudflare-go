// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package byoipprefix

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ByoIpPrefixConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/byo_ip_prefix#account_id ByoIpPrefix#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/byo_ip_prefix#asn ByoIpPrefix#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// IP Prefix in Classless Inter-Domain Routing format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/byo_ip_prefix#cidr ByoIpPrefix#cidr}
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Identifier for the uploaded LOA document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/byo_ip_prefix#loa_document_id ByoIpPrefix#loa_document_id}
	LoaDocumentId *string `field:"required" json:"loaDocumentId" yaml:"loaDocumentId"`
	// Description of the prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/byo_ip_prefix#description ByoIpPrefix#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

