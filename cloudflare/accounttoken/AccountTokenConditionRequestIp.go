// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accounttoken


type AccountTokenConditionRequestIp struct {
	// List of IPv4/IPv6 CIDR addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#in AccountToken#in}
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#not_in AccountToken#not_in}
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}

