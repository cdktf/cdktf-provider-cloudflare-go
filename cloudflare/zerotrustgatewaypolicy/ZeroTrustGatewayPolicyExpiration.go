// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyExpiration struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero
	// offsets are accepted but will be converted to the equivalent
	// value with offset zero (UTC+00:00) and will be returned as time
	// stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of
	// clients they are applied to, and expire "globally" at the point
	// given by their `expires_at` value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_gateway_policy#expires_at ZeroTrustGatewayPolicy#expires_at}
	ExpiresAt *string `field:"required" json:"expiresAt" yaml:"expiresAt"`
	// The default duration a policy will be active in minutes.
	//
	// Must be set in order to use the `reset_expiration` endpoint on this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_gateway_policy#duration ZeroTrustGatewayPolicy#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
}

