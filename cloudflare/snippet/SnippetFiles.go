// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snippet


type SnippetFiles struct {
	// Name of the snippet file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/snippet#name Snippet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Content of the snippet file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/snippet#content Snippet#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
}

