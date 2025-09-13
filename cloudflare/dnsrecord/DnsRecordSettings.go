// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecord


type DnsRecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting address records (e.g., A and AAAA) to be returned instead of the CNAME record itself. This setting is unavailable for proxied records, since they are always flattened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/dns_record#flatten_cname DnsRecord#flatten_cname}
	FlattenCname interface{} `field:"optional" json:"flattenCname" yaml:"flattenCname"`
	// When enabled, only A records will be generated, and AAAA records will not be created.
	//
	// This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/dns_record#ipv4_only DnsRecord#ipv4_only}
	Ipv4Only interface{} `field:"optional" json:"ipv4Only" yaml:"ipv4Only"`
	// When enabled, only AAAA records will be generated, and A records will not be created.
	//
	// This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/dns_record#ipv6_only DnsRecord#ipv6_only}
	Ipv6Only interface{} `field:"optional" json:"ipv6Only" yaml:"ipv6Only"`
}

