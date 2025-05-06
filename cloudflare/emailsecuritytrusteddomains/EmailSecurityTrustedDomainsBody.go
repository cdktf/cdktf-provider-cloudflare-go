// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailsecuritytrusteddomains


type EmailSecurityTrustedDomainsBody struct {
	// Select to prevent recently registered domains from triggering a Suspicious or Malicious disposition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/email_security_trusted_domains#is_recent EmailSecurityTrustedDomains#is_recent}
	IsRecent interface{} `field:"required" json:"isRecent" yaml:"isRecent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/email_security_trusted_domains#is_regex EmailSecurityTrustedDomains#is_regex}.
	IsRegex interface{} `field:"required" json:"isRegex" yaml:"isRegex"`
	// Select for partner or other approved domains that have similar spelling to your connected domains.
	//
	// Prevents listed domains from
	// triggering a Spoof disposition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/email_security_trusted_domains#is_similarity EmailSecurityTrustedDomains#is_similarity}
	IsSimilarity interface{} `field:"required" json:"isSimilarity" yaml:"isSimilarity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/email_security_trusted_domains#pattern EmailSecurityTrustedDomains#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/email_security_trusted_domains#comments EmailSecurityTrustedDomains#comments}.
	Comments *string `field:"optional" json:"comments" yaml:"comments"`
}

