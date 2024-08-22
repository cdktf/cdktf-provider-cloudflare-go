// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyIncludeExternalEvaluation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_policy#evaluate_url ZeroTrustAccessPolicy#evaluate_url}.
	EvaluateUrl *string `field:"optional" json:"evaluateUrl" yaml:"evaluateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.40.0/docs/resources/zero_trust_access_policy#keys_url ZeroTrustAccessPolicy#keys_url}.
	KeysUrl *string `field:"optional" json:"keysUrl" yaml:"keysUrl"`
}

