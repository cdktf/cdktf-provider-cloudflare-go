// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list


type ListItem struct {
	// An optional comment for the item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/list#comment List#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/list#value List#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

