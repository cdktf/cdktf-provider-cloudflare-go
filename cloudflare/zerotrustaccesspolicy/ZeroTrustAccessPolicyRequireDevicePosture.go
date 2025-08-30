// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy


type ZeroTrustAccessPolicyRequireDevicePosture struct {
	// The ID of a device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/zero_trust_access_policy#integration_uid ZeroTrustAccessPolicy#integration_uid}
	IntegrationUid *string `field:"required" json:"integrationUid" yaml:"integrationUid"`
}

