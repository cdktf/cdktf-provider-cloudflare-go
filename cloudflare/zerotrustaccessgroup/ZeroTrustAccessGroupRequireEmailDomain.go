// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessgroup


type ZeroTrustAccessGroupRequireEmailDomain struct {
	// The email domain to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_group#domain ZeroTrustAccessGroup#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

