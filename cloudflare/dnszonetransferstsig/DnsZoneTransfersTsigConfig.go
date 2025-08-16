// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnszonetransferstsig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneTransfersTsigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/dns_zone_transfers_tsig#account_id DnsZoneTransfersTsig#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// TSIG algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/dns_zone_transfers_tsig#algo DnsZoneTransfersTsig#algo}
	Algo *string `field:"required" json:"algo" yaml:"algo"`
	// TSIG key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/dns_zone_transfers_tsig#name DnsZoneTransfersTsig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// TSIG secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/dns_zone_transfers_tsig#secret DnsZoneTransfersTsig#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
}

