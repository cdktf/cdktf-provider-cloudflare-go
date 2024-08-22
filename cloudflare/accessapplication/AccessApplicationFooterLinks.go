// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationFooterLinks struct {
	// The name of the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL of the footer link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/access_application#url AccessApplication#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

