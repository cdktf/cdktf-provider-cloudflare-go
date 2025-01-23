// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package leakedcredentialcheckrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LeakedCredentialCheckRuleConfig struct {
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
	// The ruleset expression to use in matching the password in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/leaked_credential_check_rule#password LeakedCredentialCheckRule#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The ruleset expression to use in matching the username in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/leaked_credential_check_rule#username LeakedCredentialCheckRule#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The zone identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.51.0/docs/resources/leaked_credential_check_rule#zone_id LeakedCredentialCheckRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

