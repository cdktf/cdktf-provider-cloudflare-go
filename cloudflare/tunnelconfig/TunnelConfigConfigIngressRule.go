// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tunnelconfig


type TunnelConfigConfigIngressRule struct {
	// Name of the service to which the request will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/tunnel_config#service TunnelConfigA#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Hostname to match the incoming request with. If the hostname matches, the request will be sent to the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/tunnel_config#hostname TunnelConfigA#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// origin_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/tunnel_config#origin_request TunnelConfigA#origin_request}
	OriginRequest *TunnelConfigConfigIngressRuleOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// Path of the incoming request. If the path matches, the request will be sent to the local service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.24.0/docs/resources/tunnel_config#path TunnelConfigA#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

