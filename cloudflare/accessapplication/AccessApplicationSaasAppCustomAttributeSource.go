// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationSaasAppCustomAttributeSource struct {
	// The name of the attribute as provided by the IDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.17.0/docs/resources/access_application#name AccessApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

