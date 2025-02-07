// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessmtlshostnamesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessMtlsHostnameSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_mtls_hostname_settings#settings ZeroTrustAccessMtlsHostnameSettings#settings}.
	Settings interface{} `field:"required" json:"settings" yaml:"settings"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_mtls_hostname_settings#account_id ZeroTrustAccessMtlsHostnameSettings#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/zero_trust_access_mtls_hostname_settings#zone_id ZeroTrustAccessMtlsHostnameSettings#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

