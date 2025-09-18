// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apitoken


type ApiTokenPolicies struct {
	// Allow or deny operations against the resources. Available values: "allow", "deny".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/api_token#effect ApiToken#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// A set of permission groups that are specified to the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/api_token#permission_groups ApiToken#permission_groups}
	PermissionGroups interface{} `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// A list of resource names that the policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/api_token#resources ApiToken#resources}
	Resources *map[string]*string `field:"required" json:"resources" yaml:"resources"`
}

