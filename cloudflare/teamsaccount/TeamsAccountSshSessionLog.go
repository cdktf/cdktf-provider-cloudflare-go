// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount


type TeamsAccountSshSessionLog struct {
	// Public key used to encrypt ssh session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.22.0/docs/resources/teams_account#public_key TeamsAccount#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

