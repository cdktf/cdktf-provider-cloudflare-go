// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package addressmap


type AddressMapMemberships struct {
	// Identifier of the account or zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/address_map#identifier AddressMap#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The type of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/address_map#kind AddressMap#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

