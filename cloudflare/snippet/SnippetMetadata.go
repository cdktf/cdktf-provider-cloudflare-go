// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippet


type SnippetMetadata struct {
	// Name of the file that contains the main module of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/snippet#main_module Snippet#main_module}
	MainModule *string `field:"required" json:"mainModule" yaml:"mainModule"`
}

