// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofilelocaldomainfallback


type ZeroTrustDeviceDefaultProfileLocalDomainFallbackDomains struct {
	// The domain suffix to match when resolving locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_device_default_profile_local_domain_fallback#suffix ZeroTrustDeviceDefaultProfileLocalDomainFallback#suffix}
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
	// A description of the fallback domain, displayed in the client UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_device_default_profile_local_domain_fallback#description ZeroTrustDeviceDefaultProfileLocalDomainFallback#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of IP addresses to handle domain resolution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/zero_trust_device_default_profile_local_domain_fallback#dns_server ZeroTrustDeviceDefaultProfileLocalDomainFallback#dns_server}
	DnsServer *[]*string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
}

