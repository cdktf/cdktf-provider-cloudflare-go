// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustinfrastructureaccesstargets


type DataCloudflareZeroTrustInfrastructureAccessTargetsTargetsIpIpv4 struct {
	// The IP address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/data-sources/zero_trust_infrastructure_access_targets#ip_addr DataCloudflareZeroTrustInfrastructureAccessTargets#ip_addr}
	IpAddr *string `field:"required" json:"ipAddr" yaml:"ipAddr"`
	// The private virtual network identifier for the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/data-sources/zero_trust_infrastructure_access_targets#virtual_network_id DataCloudflareZeroTrustInfrastructureAccessTargets#virtual_network_id}
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

