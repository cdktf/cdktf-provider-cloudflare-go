// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessapplication


type AccessApplicationDestinations struct {
	// The private CIDR of the destination.
	//
	// Only valid when type=private. IPs are computed as /32 cidr. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#cidr AccessApplication#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The private hostname of the destination.
	//
	// Only valid when type=private. Private hostnames currently match only Server Name Indications (SNI). Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#hostname AccessApplication#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The l4 protocol that matches this destination.
	//
	// Only valid when type=private. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#l4_protocol AccessApplication#l4_protocol}
	L4Protocol *string `field:"optional" json:"l4Protocol" yaml:"l4Protocol"`
	// The port range of the destination.
	//
	// Only valid when type=private. Single ports are supported. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#port_range AccessApplication#port_range}
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// The destination type. Available values: `public`, `private`. Defaults to `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#type AccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The public URI of the destination. Can include a domain and path with wildcards. Only valid when type=public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#uri AccessApplication#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The VNet ID of the destination.
	//
	// Only valid when type=private. Private destinations are an early access feature and gated behind a feature flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/access_application#vnet_id AccessApplication#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

