// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsaccount


type TeamsAccountProxy struct {
	// Sets the time limit in seconds that a user can use an override code to bypass WARP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#disable_for_time TeamsAccount#disable_for_time}
	DisableForTime *float64 `field:"required" json:"disableForTime" yaml:"disableForTime"`
	// Whether root ca is enabled account wide for ZT clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#root_ca TeamsAccount#root_ca}
	RootCa interface{} `field:"required" json:"rootCa" yaml:"rootCa"`
	// Whether gateway proxy is enabled on gateway devices for TCP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#tcp TeamsAccount#tcp}
	Tcp interface{} `field:"required" json:"tcp" yaml:"tcp"`
	// Whether gateway proxy is enabled on gateway devices for UDP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#udp TeamsAccount#udp}
	Udp interface{} `field:"required" json:"udp" yaml:"udp"`
	// Whether virtual IP (CGNAT) is enabled account wide and will override existing local interface IP for ZT clients.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/teams_account#virtual_ip TeamsAccount#virtual_ip}
	VirtualIp interface{} `field:"required" json:"virtualIp" yaml:"virtualIp"`
}

