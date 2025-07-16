// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdnslocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDnsLocationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#account_id ZeroTrustDnsLocation#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#name ZeroTrustDnsLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// True if the location is the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#client_default ZeroTrustDnsLocation#client_default}
	ClientDefault interface{} `field:"optional" json:"clientDefault" yaml:"clientDefault"`
	// The identifier of the pair of IPv4 addresses assigned to this location.
	//
	// When creating a location, if this field is absent or set with null, the pair of shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned. When updating a location, if the field is absent or set with null, the pre-assigned pair remains unchanged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#dns_destination_ips_id ZeroTrustDnsLocation#dns_destination_ips_id}
	DnsDestinationIpsId *string `field:"optional" json:"dnsDestinationIpsId" yaml:"dnsDestinationIpsId"`
	// True if the location needs to resolve EDNS queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#ecs_support ZeroTrustDnsLocation#ecs_support}
	EcsSupport interface{} `field:"optional" json:"ecsSupport" yaml:"ecsSupport"`
	// The destination endpoints configured for this location.
	//
	// When updating a location, if this field is absent or set with null, the endpoints configuration remains unchanged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#endpoints ZeroTrustDnsLocation#endpoints}
	Endpoints *ZeroTrustDnsLocationEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// A list of network ranges that requests from this location would originate from.
	//
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dns_location#networks ZeroTrustDnsLocation#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

