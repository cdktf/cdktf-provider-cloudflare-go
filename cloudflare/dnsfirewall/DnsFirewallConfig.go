// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsfirewall

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsFirewallConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#account_id DnsFirewall#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// DNS Firewall cluster name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#name DnsFirewall#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#upstream_ips DnsFirewall#upstream_ips}.
	UpstreamIps *[]*string `field:"required" json:"upstreamIps" yaml:"upstreamIps"`
	// Attack mitigation settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#attack_mitigation DnsFirewall#attack_mitigation}
	AttackMitigation *DnsFirewallAttackMitigation `field:"optional" json:"attackMitigation" yaml:"attackMitigation"`
	// Whether to refuse to answer queries for the ANY type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#deprecate_any_requests DnsFirewall#deprecate_any_requests}
	DeprecateAnyRequests interface{} `field:"optional" json:"deprecateAnyRequests" yaml:"deprecateAnyRequests"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#ecs_fallback DnsFirewall#ecs_fallback}
	EcsFallback interface{} `field:"optional" json:"ecsFallback" yaml:"ecsFallback"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers.
	//
	// Higher TTLs will be decreased to the maximum defined here for caching purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#maximum_cache_ttl DnsFirewall#maximum_cache_ttl}
	MaximumCacheTtl *float64 `field:"optional" json:"maximumCacheTtl" yaml:"maximumCacheTtl"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes of caching between DNS Firewall and the upstream servers.
	//
	// Lower TTLs will be increased to the minimum defined here for caching purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#minimum_cache_ttl DnsFirewall#minimum_cache_ttl}
	MinimumCacheTtl *float64 `field:"optional" json:"minimumCacheTtl" yaml:"minimumCacheTtl"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache negative responses (e.g., NXDOMAIN) from the upstream servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#negative_cache_ttl DnsFirewall#negative_cache_ttl}
	NegativeCacheTtl *float64 `field:"optional" json:"negativeCacheTtl" yaml:"negativeCacheTtl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to the upstream nameservers configured on the cluster).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#ratelimit DnsFirewall#ratelimit}
	Ratelimit *float64 `field:"optional" json:"ratelimit" yaml:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not counting the initial attempt).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/dns_firewall#retries DnsFirewall#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
}

