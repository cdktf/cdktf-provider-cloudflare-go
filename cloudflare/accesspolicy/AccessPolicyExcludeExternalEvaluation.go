// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesspolicy


type AccessPolicyExcludeExternalEvaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.31.0/docs/resources/access_policy#evaluate_url AccessPolicy#evaluate_url}.
	EvaluateUrl *string `field:"optional" json:"evaluateUrl" yaml:"evaluateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.31.0/docs/resources/access_policy#keys_url AccessPolicy#keys_url}.
	KeysUrl *string `field:"optional" json:"keysUrl" yaml:"keysUrl"`
}

