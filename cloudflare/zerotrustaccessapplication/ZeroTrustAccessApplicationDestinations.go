// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationDestinations struct {
	// The private CIDR of the destination.
	//
	// Only valid when type=private. IPs are computed as /32 cidr. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#cidr ZeroTrustAccessApplication#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The private hostname of the destination.
	//
	// Only valid when type=private. Private hostnames currently match only Server Name Indications (SNI). Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#hostname ZeroTrustAccessApplication#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The l4 protocol that matches this destination.
	//
	// Only valid when type=private. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#l4_protocol ZeroTrustAccessApplication#l4_protocol}
	L4Protocol *string `field:"optional" json:"l4Protocol" yaml:"l4Protocol"`
	// The port range of the destination.
	//
	// Only valid when type=private. Single ports are supported. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#port_range ZeroTrustAccessApplication#port_range}
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// The destination type. Available values: `public`, `private`. Defaults to `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#type ZeroTrustAccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The public URI of the destination. Can include a domain and path with wildcards. Only valid when type=public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#uri ZeroTrustAccessApplication#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The VNet ID of the destination.
	//
	// Only valid when type=private. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_application#vnet_id ZeroTrustAccessApplication#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

