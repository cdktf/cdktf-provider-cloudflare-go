// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationScimConfigMappingsOperations struct {
	// Whether or not this mapping applies to create (POST) operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_application#create ZeroTrustAccessApplication#create}
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// Whether or not this mapping applies to DELETE operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_application#delete ZeroTrustAccessApplication#delete}
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Whether or not this mapping applies to update (PATCH/PUT) operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_application#update ZeroTrustAccessApplication#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

