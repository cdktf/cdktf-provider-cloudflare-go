// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailsecuritytrusteddomains

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailSecurityTrustedDomainsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Account Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#account_id EmailSecurityTrustedDomains#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#body EmailSecurityTrustedDomains#body}.
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#comments EmailSecurityTrustedDomains#comments}.
	Comments *string `field:"optional" json:"comments" yaml:"comments"`
	// Select to prevent recently registered domains from triggering a Suspicious or Malicious disposition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#is_recent EmailSecurityTrustedDomains#is_recent}
	IsRecent interface{} `field:"optional" json:"isRecent" yaml:"isRecent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#is_regex EmailSecurityTrustedDomains#is_regex}.
	IsRegex interface{} `field:"optional" json:"isRegex" yaml:"isRegex"`
	// Select for partner or other approved domains that have similar spelling to your connected domains.
	//
	// Prevents listed domains from
	// triggering a Spoof disposition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#is_similarity EmailSecurityTrustedDomains#is_similarity}
	IsSimilarity interface{} `field:"optional" json:"isSimilarity" yaml:"isSimilarity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/email_security_trusted_domains#pattern EmailSecurityTrustedDomains#pattern}.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

