// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnszonetransferspeer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneTransfersPeerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#account_id DnsZoneTransfersPeer#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the peer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#name DnsZoneTransfersPeer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone this peer is linked to.
	//
	// For primary zones this IP defines the IP of the secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary zones this IP defines the IP of the primary nameserver Cloudflare will send AXFR/IXFR requests to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#ip DnsZoneTransfersPeer#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#ixfr_enable DnsZoneTransfersPeer#ixfr_enable}
	IxfrEnable interface{} `field:"optional" json:"ixfrEnable" yaml:"ixfrEnable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is linked to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#port DnsZoneTransfersPeer#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_zone_transfers_peer#tsig_id DnsZoneTransfersPeer#tsig_id}
	TsigId *string `field:"optional" json:"tsigId" yaml:"tsigId"`
}

