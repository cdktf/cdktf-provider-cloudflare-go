// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/ruleset#public_key Ruleset#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

