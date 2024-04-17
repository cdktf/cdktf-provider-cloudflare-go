// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamslocation


type TeamsLocationNetworks struct {
	// CIDR notation representation of the network IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.30.0/docs/resources/teams_location#network TeamsLocation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

