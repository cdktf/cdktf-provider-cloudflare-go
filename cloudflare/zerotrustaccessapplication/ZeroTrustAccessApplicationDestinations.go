// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationDestinations struct {
	// The CIDR range of the destination. Single IPs will be computed as /32.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#cidr ZeroTrustAccessApplication#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The hostname of the destination. Matches a valid SNI served by an HTTPS origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#hostname ZeroTrustAccessApplication#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The L4 protocol of the destination. When omitted, both UDP and TCP traffic will match. Available values: "tcp", "udp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#l4_protocol ZeroTrustAccessApplication#l4_protocol}
	L4Protocol *string `field:"optional" json:"l4Protocol" yaml:"l4Protocol"`
	// The port range of the destination.
	//
	// Can be a single port or a range of ports. When omitted, all ports will match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#port_range ZeroTrustAccessApplication#port_range}
	PortRange *string `field:"optional" json:"portRange" yaml:"portRange"`
	// Available values: "public", "private".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#type ZeroTrustAccessApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The URI of the destination. Public destinations' URIs can include a domain and path with [wildcards](https://developers.cloudflare.com/cloudflare-one/policies/access/app-paths/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#uri ZeroTrustAccessApplication#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
	// The VNET ID to match the destination. When omitted, all VNETs will match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/zero_trust_access_application#vnet_id ZeroTrustAccessApplication#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

