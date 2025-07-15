// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitconnector


type MagicTransitConnectorDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/magic_transit_connector#id MagicTransitConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/magic_transit_connector#serial_number MagicTransitConnector#serial_number}.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
}

