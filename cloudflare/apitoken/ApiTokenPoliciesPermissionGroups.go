// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apitoken


type ApiTokenPoliciesPermissionGroups struct {
	// Identifier of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/api_token#id ApiToken#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Attributes associated to the permission group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/api_token#meta ApiToken#meta}
	Meta *ApiTokenPoliciesPermissionGroupsMeta `field:"optional" json:"meta" yaml:"meta"`
}

