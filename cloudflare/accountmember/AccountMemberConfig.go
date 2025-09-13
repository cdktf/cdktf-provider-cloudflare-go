// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountMemberConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account_member#account_id AccountMember#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The contact email address of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account_member#email AccountMember#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Array of policies associated with this member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account_member#policies AccountMember#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// Array of roles associated with this member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account_member#roles AccountMember#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Available values: "accepted", "pending".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/account_member#status AccountMember#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

