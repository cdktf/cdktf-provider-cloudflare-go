// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicecustomprofilelocaldomainfallback


type ZeroTrustDeviceCustomProfileLocalDomainFallbackDomains struct {
	// The domain suffix to match when resolving locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#suffix ZeroTrustDeviceCustomProfileLocalDomainFallback#suffix}
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
	// A description of the fallback domain, displayed in the client UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#description ZeroTrustDeviceCustomProfileLocalDomainFallback#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of IP addresses to handle domain resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_device_custom_profile_local_domain_fallback#dns_server ZeroTrustDeviceCustomProfileLocalDomainFallback#dns_server}
	DnsServer *[]*string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
}

