// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessmutualtlshostnamesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessMutualTlsHostnameSettingsConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#account_id AccessMutualTlsHostnameSettings#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#settings AccessMutualTlsHostnameSettings#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#zone_id AccessMutualTlsHostnameSettings#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

