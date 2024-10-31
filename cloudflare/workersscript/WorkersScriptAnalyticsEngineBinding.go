// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptAnalyticsEngineBinding struct {
	// The name of the Analytics Engine dataset to write to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/workers_script#dataset WorkersScript#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

