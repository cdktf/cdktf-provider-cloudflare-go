// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FilterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#body Filter#body}.
	Body interface{} `field:"required" json:"body" yaml:"body"`
	// Defines an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#zone_id Filter#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// An informative summary of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#description Filter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#expression Filter#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// When true, indicates that the filter is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#paused Filter#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// A short reference tag. Allows you to select related filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/filter#ref Filter#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

