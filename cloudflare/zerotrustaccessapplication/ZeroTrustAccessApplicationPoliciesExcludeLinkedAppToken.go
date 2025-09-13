// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesExcludeLinkedAppToken struct {
	// The ID of an Access OIDC SaaS application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#app_uid ZeroTrustAccessApplication#app_uid}
	AppUid *string `field:"required" json:"appUid" yaml:"appUid"`
}

