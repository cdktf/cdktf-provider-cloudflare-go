// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsProductionServiceBinding struct {
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/pages_project#name PagesProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/pages_project#service PagesProject#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// The name of the Worker environment to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/pages_project#environment PagesProject#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
}

