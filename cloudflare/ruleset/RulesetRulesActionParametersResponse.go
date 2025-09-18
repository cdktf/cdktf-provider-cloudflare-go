// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParametersResponse struct {
	// The content to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#content Ruleset#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The type of the content to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The status code to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"required" json:"statusCode" yaml:"statusCode"`
}

