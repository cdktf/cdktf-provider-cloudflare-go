// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamslocation


type TeamsLocationEndpoints struct {
	// doh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_location#doh TeamsLocation#doh}
	Doh *TeamsLocationEndpointsDoh `field:"optional" json:"doh" yaml:"doh"`
	// dot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_location#dot TeamsLocation#dot}
	Dot *TeamsLocationEndpointsDot `field:"optional" json:"dot" yaml:"dot"`
	// ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_location#ipv4 TeamsLocation#ipv4}
	Ipv4 *TeamsLocationEndpointsIpv4 `field:"optional" json:"ipv4" yaml:"ipv4"`
	// ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/teams_location#ipv6 TeamsLocation#ipv6}
	Ipv6 *TeamsLocationEndpointsIpv6 `field:"optional" json:"ipv6" yaml:"ipv6"`
}

