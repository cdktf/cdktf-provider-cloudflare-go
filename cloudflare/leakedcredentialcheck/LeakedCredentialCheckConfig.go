// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package leakedcredentialcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LeakedCredentialCheckConfig struct {
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
	// State of the Leaked Credential Check detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/leaked_credential_check#enabled LeakedCredentialCheck#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/leaked_credential_check#zone_id LeakedCredentialCheck#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

