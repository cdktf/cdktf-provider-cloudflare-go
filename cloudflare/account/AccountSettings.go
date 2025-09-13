// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account


type AccountSettings struct {
	// Sets an abuse contact email to notify for abuse reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account#abuse_contact_email Account#abuse_contact_email}
	AbuseContactEmail *string `field:"optional" json:"abuseContactEmail" yaml:"abuseContactEmail"`
	// Indicates whether membership in this account requires that Two-Factor Authentication is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account#enforce_twofactor Account#enforce_twofactor}
	EnforceTwofactor interface{} `field:"optional" json:"enforceTwofactor" yaml:"enforceTwofactor"`
}

