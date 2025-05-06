// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accounttoken


type AccountTokenCondition struct {
	// Client IP restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/account_token#request_ip AccountToken#request_ip}
	RequestIp *AccountTokenConditionRequestIp `field:"optional" json:"requestIp" yaml:"requestIp"`
}

