// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyConnectionRules struct {
	// ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/access_policy#ssh AccessPolicy#ssh}
	Ssh *AccessPolicyConnectionRulesSsh `field:"required" json:"ssh" yaml:"ssh"`
}

