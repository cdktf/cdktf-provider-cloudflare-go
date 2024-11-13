// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareinfrastructureaccesstargets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareInfrastructureAccessTargetsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#account_id DataCloudflareInfrastructureAccessTargets#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A date and time after a target was created to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#created_after DataCloudflareInfrastructureAccessTargets#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// The hostname of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#hostname DataCloudflareInfrastructureAccessTargets#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Partial match to the hostname of a target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#hostname_contains DataCloudflareInfrastructureAccessTargets#hostname_contains}
	HostnameContains *string `field:"optional" json:"hostnameContains" yaml:"hostnameContains"`
	// The target's IPv4 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#ipv4 DataCloudflareInfrastructureAccessTargets#ipv4}
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// The target's IPv6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#ipv6 DataCloudflareInfrastructureAccessTargets#ipv6}
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// A date and time after a target was modified to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#modified_after DataCloudflareInfrastructureAccessTargets#modified_after}
	ModifiedAfter *string `field:"optional" json:"modifiedAfter" yaml:"modifiedAfter"`
	// The private virtual network identifier for the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.46.0/docs/data-sources/infrastructure_access_targets#virtual_network_id DataCloudflareInfrastructureAccessTargets#virtual_network_id}
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

