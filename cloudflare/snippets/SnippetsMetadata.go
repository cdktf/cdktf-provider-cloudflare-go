// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippets


type SnippetsMetadata struct {
	// Main module name of uploaded snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/snippets#main_module Snippets#main_module}
	MainModule *string `field:"optional" json:"mainModule" yaml:"mainModule"`
}

