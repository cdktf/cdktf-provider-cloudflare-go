// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apitoken


type ApiTokenConditionRequestIp struct {
	// List of IPv4/IPv6 CIDR addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/api_token#in ApiToken#in}
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/api_token#not_in ApiToken#not_in}
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}

