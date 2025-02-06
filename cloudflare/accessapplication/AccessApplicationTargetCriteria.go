// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationTargetCriteria struct {
	// The port that the targets use for the chosen communication protocol. A port cannot be assigned to multiple protocols.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#port AccessApplication#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The communication protocol your application secures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#protocol AccessApplication#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// target_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#target_attributes AccessApplication#target_attributes}
	TargetAttributes interface{} `field:"required" json:"targetAttributes" yaml:"targetAttributes"`
}

