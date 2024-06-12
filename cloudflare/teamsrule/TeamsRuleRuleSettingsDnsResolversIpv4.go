// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsrule


type TeamsRuleRuleSettingsDnsResolversIpv4 struct {
	// The IPv4 or IPv6 address of the upstream resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/teams_rule#ip TeamsRule#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// A port number to use for the upstream resolver. Defaults to `53`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/teams_rule#port TeamsRule#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Whether to connect to this resolver over a private network. Must be set when `vnet_id` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/teams_rule#route_through_private_network TeamsRule#route_through_private_network}
	RouteThroughPrivateNetwork interface{} `field:"optional" json:"routeThroughPrivateNetwork" yaml:"routeThroughPrivateNetwork"`
	// specify a virtual network for this resolver. Uses default virtual network id if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.35.0/docs/resources/teams_rule#vnet_id TeamsRule#vnet_id}
	VnetId *string `field:"optional" json:"vnetId" yaml:"vnetId"`
}

