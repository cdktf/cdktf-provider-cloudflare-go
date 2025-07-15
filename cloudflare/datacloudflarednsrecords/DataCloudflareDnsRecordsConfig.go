// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarednsrecords

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareDnsRecordsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#zone_id DataCloudflareDnsRecords#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#comment DataCloudflareDnsRecords#comment}.
	Comment *DataCloudflareDnsRecordsComment `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#content DataCloudflareDnsRecords#content}.
	Content *DataCloudflareDnsRecordsContent `field:"optional" json:"content" yaml:"content"`
	// Direction to order DNS records in. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#direction DataCloudflareDnsRecords#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Whether to match all search requirements or at least one (any).
	//
	// If set to `all`, acts like a logical AND between filters. If set to `any`, acts like a logical OR instead. Note that the interaction between tag filters is controlled by the `tag-match` parameter instead.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#match DataCloudflareDnsRecords#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#max_items DataCloudflareDnsRecords#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#name DataCloudflareDnsRecords#name}.
	Name *DataCloudflareDnsRecordsName `field:"optional" json:"name" yaml:"name"`
	// Field to order DNS records by. Available values: "type", "name", "content", "ttl", "proxied".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#order DataCloudflareDnsRecords#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Whether the record is receiving the performance and security benefits of Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#proxied DataCloudflareDnsRecords#proxied}
	Proxied interface{} `field:"optional" json:"proxied" yaml:"proxied"`
	// Allows searching in multiple properties of a DNS record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact behavior is intentionally left unspecified and is subject to change in the future. This parameter works independently of the `match` setting. For automated searches, please use the other available parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#search DataCloudflareDnsRecords#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#tag DataCloudflareDnsRecords#tag}.
	Tag *DataCloudflareDnsRecordsTag `field:"optional" json:"tag" yaml:"tag"`
	// Whether to match all tag search requirements or at least one (any).
	//
	// If set to `all`, acts like a logical AND between tag filters. If set to `any`, acts like a logical OR instead. Note that the regular `match` parameter is still used to combine the resulting condition with other filters that aren't related to tags.
	// Available values: "any", "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#tag_match DataCloudflareDnsRecords#tag_match}
	TagMatch *string `field:"optional" json:"tagMatch" yaml:"tagMatch"`
	// Record type.
	//
	// Available values: "A", "AAAA", "CAA", "CERT", "CNAME", "DNSKEY", "DS", "HTTPS", "LOC", "MX", "NAPTR", "NS", "OPENPGPKEY", "PTR", "SMIMEA", "SRV", "SSHFP", "SVCB", "TLSA", "TXT", "URI".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/dns_records#type DataCloudflareDnsRecords#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

