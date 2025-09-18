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
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/leaked_credential_check_rule#zone_id LeakedCredentialCheckRule#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/leaked_credential_check_rule#password LeakedCredentialCheckRule#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/leaked_credential_check_rule#username LeakedCredentialCheckRule#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

