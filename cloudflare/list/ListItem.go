// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package list


type ListItem struct {
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/list#value List#value}
	Value *ListItemValue `field:"required" json:"value" yaml:"value"`
	// An optional comment for the item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/list#comment List#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

