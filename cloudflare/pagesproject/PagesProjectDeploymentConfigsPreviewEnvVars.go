// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsPreviewEnvVars struct {
	// Available values: "plain_text", "secret_text".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/pages_project#type PagesProject#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Environment variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/pages_project#value PagesProject#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

