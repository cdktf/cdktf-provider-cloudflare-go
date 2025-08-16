// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spectrumapplication


type SpectrumApplicationEdgeIps struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs. Available values: "all", "ipv4", "ipv6".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/spectrum_application#connectivity SpectrumApplication#connectivity}
	Connectivity *string `field:"optional" json:"connectivity" yaml:"connectivity"`
	// The array of customer owned IPs we broadcast via anycast for this hostname and application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/spectrum_application#ips SpectrumApplication#ips}
	Ips *[]*string `field:"optional" json:"ips" yaml:"ips"`
	// The type of edge IP configuration specified.
	//
	// Dynamically allocated edge IPs use Spectrum anycast IPs in accordance with the connectivity you specify. Only valid with CNAME DNS names.
	// Available values: "dynamic", "static".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/spectrum_application#type SpectrumApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

