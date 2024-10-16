// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationScimConfigMappingsOperations struct {
	// Whether or not this mapping applies to create (POST) operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/access_application#create AccessApplication#create}
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// Whether or not this mapping applies to DELETE operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/access_application#delete AccessApplication#delete}
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.44.0/docs/resources/access_application#update AccessApplication#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

