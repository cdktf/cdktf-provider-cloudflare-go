// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsProductionPlacement struct {
	// Placement Mode for the Pages Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/pages_project#mode PagesProject#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

