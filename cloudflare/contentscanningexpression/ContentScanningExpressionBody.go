// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contentscanningexpression


type ContentScanningExpressionBody struct {
	// Defines the ruleset expression to use in matching content objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/content_scanning_expression#payload ContentScanningExpression#payload}
	Payload *string `field:"required" json:"payload" yaml:"payload"`
}

