// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptD1DatabaseBinding struct {
	// Database ID of D1 database to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/workers_script#database_id WorkersScript#database_id}
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

