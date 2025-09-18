// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsProductionServices struct {
	// The entrypoint to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#entrypoint PagesProject#entrypoint}
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The Service environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#environment PagesProject#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The Service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/pages_project#service PagesProject#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

