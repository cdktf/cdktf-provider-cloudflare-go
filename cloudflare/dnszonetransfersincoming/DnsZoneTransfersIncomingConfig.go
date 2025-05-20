// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnszonetransfersincoming

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneTransfersIncomingConfig struct {
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
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not applicable for primary zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/dns_zone_transfers_incoming#auto_refresh_seconds DnsZoneTransfersIncoming#auto_refresh_seconds}
	AutoRefreshSeconds *float64 `field:"required" json:"autoRefreshSeconds" yaml:"autoRefreshSeconds"`
	// Zone name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/dns_zone_transfers_incoming#name DnsZoneTransfersIncoming#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of peer tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/dns_zone_transfers_incoming#peers DnsZoneTransfersIncoming#peers}
	Peers *[]*string `field:"required" json:"peers" yaml:"peers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/dns_zone_transfers_incoming#zone_id DnsZoneTransfersIncoming#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

