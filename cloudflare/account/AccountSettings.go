// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package account


type AccountSettings struct {
	// Sets an abuse contact email to notify for abuse reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/account#abuse_contact_email Account#abuse_contact_email}
	AbuseContactEmail *string `field:"optional" json:"abuseContactEmail" yaml:"abuseContactEmail"`
	// Specifies the default nameservers to be used for new zones added to this account.
	//
	// - `cloudflare.standard` for Cloudflare-branded nameservers
	// - `custom.account` for account custom nameservers
	// - `custom.tenant` for tenant custom nameservers
	//
	// See [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
	// for more information.
	//
	// Deprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	// Available values: "cloudflare.standard", "custom.account", "custom.tenant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/account#default_nameservers Account#default_nameservers}
	DefaultNameservers *string `field:"optional" json:"defaultNameservers" yaml:"defaultNameservers"`
	// Indicates whether membership in this account requires that Two-Factor Authentication is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/account#enforce_twofactor Account#enforce_twofactor}
	EnforceTwofactor interface{} `field:"optional" json:"enforceTwofactor" yaml:"enforceTwofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by default.
	//
	// Deprecated in favor of [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/account#use_account_custom_ns_by_default Account#use_account_custom_ns_by_default}
	UseAccountCustomNsByDefault interface{} `field:"optional" json:"useAccountCustomNsByDefault" yaml:"useAccountCustomNsByDefault"`
}

