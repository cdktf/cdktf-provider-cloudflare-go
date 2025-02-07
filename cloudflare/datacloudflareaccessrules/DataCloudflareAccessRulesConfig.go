// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareaccessrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareAccessRulesConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#account_id DataCloudflareAccessRules#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#configuration DataCloudflareAccessRules#configuration}.
	Configuration *DataCloudflareAccessRulesConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// The direction used to sort returned rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#direction DataCloudflareAccessRules#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// When set to `all`, all the search requirements must match.
	//
	// When set to `any`, only one of the search requirements has to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#match DataCloudflareAccessRules#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#max_items DataCloudflareAccessRules#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// The action to apply to a matched request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#mode DataCloudflareAccessRules#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The string to search for in the notes of existing IP Access rules.
	//
	// Notes: For example, the string 'attack' would match IP Access rules with notes 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#notes DataCloudflareAccessRules#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The field used to sort returned rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#order DataCloudflareAccessRules#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/access_rules#zone_id DataCloudflareAccessRules#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

