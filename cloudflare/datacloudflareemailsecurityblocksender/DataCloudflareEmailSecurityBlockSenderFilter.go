// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareemailsecurityblocksender


type DataCloudflareEmailSecurityBlockSenderFilter struct {
	// The sorting direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_block_sender#direction DataCloudflareEmailSecurityBlockSender#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The field to sort by. Available values: "pattern", "created_at".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_block_sender#order DataCloudflareEmailSecurityBlockSender#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Available values: "EMAIL", "DOMAIN", "IP", "UNKNOWN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_block_sender#pattern_type DataCloudflareEmailSecurityBlockSender#pattern_type}
	PatternType *string `field:"optional" json:"patternType" yaml:"patternType"`
	// Allows searching in multiple properties of a record simultaneously.
	//
	// This parameter is intended for human users, not automation. Its exact
	// behavior is intentionally left unspecified and is subject to change
	// in the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/email_security_block_sender#search DataCloudflareEmailSecurityBlockSender#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

