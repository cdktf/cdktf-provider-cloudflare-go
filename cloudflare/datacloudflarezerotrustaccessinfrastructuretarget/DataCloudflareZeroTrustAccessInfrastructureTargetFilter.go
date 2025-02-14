// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarezerotrustaccessinfrastructuretarget


type DataCloudflareZeroTrustAccessInfrastructureTargetFilter struct {
	// Date and time at which the target was created after (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#created_after DataCloudflareZeroTrustAccessInfrastructureTarget#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Date and time at which the target was created before (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#created_before DataCloudflareZeroTrustAccessInfrastructureTarget#created_before}
	CreatedBefore *string `field:"optional" json:"createdBefore" yaml:"createdBefore"`
	// The sorting direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#direction DataCloudflareZeroTrustAccessInfrastructureTarget#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Hostname of a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#hostname DataCloudflareZeroTrustAccessInfrastructureTarget#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Partial match to the hostname of a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#hostname_contains DataCloudflareZeroTrustAccessInfrastructureTarget#hostname_contains}
	HostnameContains *string `field:"optional" json:"hostnameContains" yaml:"hostnameContains"`
	// Filters for targets whose IP addresses look like the specified string. Supports `*` as a wildcard character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ip_like DataCloudflareZeroTrustAccessInfrastructureTarget#ip_like}
	IpLike *string `field:"optional" json:"ipLike" yaml:"ipLike"`
	// Filters for targets that have any of the following IP addresses.
	//
	// Specify
	// `ips` multiple times in query parameter to build list of candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ips DataCloudflareZeroTrustAccessInfrastructureTarget#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// IPv4 address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ip_v4 DataCloudflareZeroTrustAccessInfrastructureTarget#ip_v4}
	IpV4 *string `field:"optional" json:"ipV4" yaml:"ipV4"`
	// Defines an IPv4 filter range's ending value (inclusive). Requires `ipv4_start` to be specified as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ipv4_end DataCloudflareZeroTrustAccessInfrastructureTarget#ipv4_end}
	Ipv4End *string `field:"optional" json:"ipv4End" yaml:"ipv4End"`
	// Defines an IPv4 filter range's starting value (inclusive). Requires `ipv4_end` to be specified as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ipv4_start DataCloudflareZeroTrustAccessInfrastructureTarget#ipv4_start}
	Ipv4Start *string `field:"optional" json:"ipv4Start" yaml:"ipv4Start"`
	// IPv6 address of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ip_v6 DataCloudflareZeroTrustAccessInfrastructureTarget#ip_v6}
	IpV6 *string `field:"optional" json:"ipV6" yaml:"ipV6"`
	// Defines an IPv6 filter range's ending value (inclusive). Requires `ipv6_start` to be specified as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ipv6_end DataCloudflareZeroTrustAccessInfrastructureTarget#ipv6_end}
	Ipv6End *string `field:"optional" json:"ipv6End" yaml:"ipv6End"`
	// Defines an IPv6 filter range's starting value (inclusive). Requires `ipv6_end` to be specified as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#ipv6_start DataCloudflareZeroTrustAccessInfrastructureTarget#ipv6_start}
	Ipv6Start *string `field:"optional" json:"ipv6Start" yaml:"ipv6Start"`
	// Date and time at which the target was modified after (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#modified_after DataCloudflareZeroTrustAccessInfrastructureTarget#modified_after}
	ModifiedAfter *string `field:"optional" json:"modifiedAfter" yaml:"modifiedAfter"`
	// Date and time at which the target was modified before (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#modified_before DataCloudflareZeroTrustAccessInfrastructureTarget#modified_before}
	ModifiedBefore *string `field:"optional" json:"modifiedBefore" yaml:"modifiedBefore"`
	// The field to sort by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#order DataCloudflareZeroTrustAccessInfrastructureTarget#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Filters for targets that have any of the following UUIDs.
	//
	// Specify
	// `target_ids` multiple times in query parameter to build list of
	// candidates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#target_ids DataCloudflareZeroTrustAccessInfrastructureTarget#target_ids}
	TargetIds *[]*string `field:"optional" json:"targetIds" yaml:"targetIds"`
	// Private virtual network identifier of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/data-sources/zero_trust_access_infrastructure_target#virtual_network_id DataCloudflareZeroTrustAccessInfrastructureTarget#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

