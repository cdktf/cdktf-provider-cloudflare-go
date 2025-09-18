// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filter


type FilterBody struct {
	// An informative summary of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/filter#description Filter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/filter#expression Filter#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// When true, indicates that the filter is currently paused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/filter#paused Filter#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// A short reference tag. Allows you to select related filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/filter#ref Filter#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

