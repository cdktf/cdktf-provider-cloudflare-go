// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#password_expression Ruleset#password_expression}
	PasswordExpression *string `field:"required" json:"passwordExpression" yaml:"passwordExpression"`
	// Expression that selects the user ID used in the credentials check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/ruleset#username_expression Ruleset#username_expression}
	UsernameExpression *string `field:"required" json:"usernameExpression" yaml:"usernameExpression"`
}

