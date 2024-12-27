// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infrastructureaccesstarget


type InfrastructureAccessTargetIpIpv4 struct {
	// The IP address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/infrastructure_access_target#ip_addr InfrastructureAccessTarget#ip_addr}
	IpAddr *string `field:"required" json:"ipAddr" yaml:"ipAddr"`
	// The private virtual network identifier for the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/infrastructure_access_target#virtual_network_id InfrastructureAccessTarget#virtual_network_id}
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

