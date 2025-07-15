// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicemanagednetworks


type ZeroTrustDeviceManagedNetworksConfigA struct {
	// A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_managed_networks#tls_sockaddr ZeroTrustDeviceManagedNetworks#tls_sockaddr}
	TlsSockaddr *string `field:"required" json:"tlsSockaddr" yaml:"tlsSockaddr"`
	// The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr.
	//
	// If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/zero_trust_device_managed_networks#sha256 ZeroTrustDeviceManagedNetworks#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
}

