// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package record


type RecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/record#create Record#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/record#update Record#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

