// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsfirewall


type DnsFirewallAttackMitigation struct {
	// When enabled, automatically mitigate random-prefix attacks to protect upstream DNS servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/dns_firewall#enabled DnsFirewall#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/dns_firewall#only_when_upstream_unhealthy DnsFirewall#only_when_upstream_unhealthy}
	OnlyWhenUpstreamUnhealthy interface{} `field:"optional" json:"onlyWhenUpstreamUnhealthy" yaml:"onlyWhenUpstreamUnhealthy"`
}

