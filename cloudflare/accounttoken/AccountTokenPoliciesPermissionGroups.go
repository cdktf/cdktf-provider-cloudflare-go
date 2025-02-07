// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accounttoken


type AccountTokenPoliciesPermissionGroups struct {
	// Attributes associated to the permission group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/account_token#meta AccountToken#meta}
	Meta *AccountTokenPoliciesPermissionGroupsMeta `field:"optional" json:"meta" yaml:"meta"`
}

