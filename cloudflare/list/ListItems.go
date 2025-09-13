// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list


type ListItems struct {
	// A non-negative 32 bit integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#asn List#asn}
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// An informative summary of the list item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#comment List#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#hostname List#hostname}
	Hostname *ListItemsHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#ip List#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The definition of the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/list#redirect List#redirect}
	Redirect *ListItemsRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

