// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecord

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordConfig struct {
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
	// DNS record name (or @ for the zone apex) in Punycode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#name DnsRecord#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Time To Live (TTL) of the DNS record in seconds.
	//
	// Setting to 1 means 'automatic'. Value must be between 60 and 86400, with the minimum reduced to 30 for Enterprise zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#ttl DnsRecord#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// Record type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#type DnsRecord#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#zone_id DnsRecord#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Comments or notes about the DNS record. This field has no effect on DNS responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#comment DnsRecord#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A valid IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#content DnsRecord#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Components of a CAA record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#data DnsRecord#data}
	Data *DnsRecordData `field:"optional" json:"data" yaml:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#priority DnsRecord#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Whether the record is receiving the performance and security benefits of Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#proxied DnsRecord#proxied}
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// Settings for the DNS record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#settings DnsRecord#settings}
	Settings *DnsRecordSettings `field:"optional" json:"settings" yaml:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/dns_record#tags DnsRecord#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

