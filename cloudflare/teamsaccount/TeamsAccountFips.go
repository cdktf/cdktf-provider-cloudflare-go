// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount


type TeamsAccountFips struct {
	// Only allow FIPS-compliant TLS configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.21.0/docs/resources/teams_account#tls TeamsAccount#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

