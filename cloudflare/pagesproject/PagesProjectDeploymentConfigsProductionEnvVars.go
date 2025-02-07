// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsProductionEnvVars struct {
	// Environment variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/pages_project#value PagesProject#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The type of environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/pages_project#type PagesProject#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

