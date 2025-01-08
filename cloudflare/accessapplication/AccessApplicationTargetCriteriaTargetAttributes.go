// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationTargetCriteriaTargetAttributes struct {
	// The key of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_application#values AccessApplication#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

