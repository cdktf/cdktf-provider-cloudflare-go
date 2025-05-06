// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupExcludeExternalEvaluation struct {
	// The API endpoint containing your business logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_group#evaluate_url ZeroTrustAccessGroup#evaluate_url}
	EvaluateUrl *string `field:"required" json:"evaluateUrl" yaml:"evaluateUrl"`
	// The API endpoint containing the key that Access uses to verify that the response came from your API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_group#keys_url ZeroTrustAccessGroup#keys_url}
	KeysUrl *string `field:"required" json:"keysUrl" yaml:"keysUrl"`
}

