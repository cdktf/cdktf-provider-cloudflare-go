// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyExcludeGeo struct {
	// The country code that should be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_access_policy#country_code ZeroTrustAccessPolicy#country_code}
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
}

