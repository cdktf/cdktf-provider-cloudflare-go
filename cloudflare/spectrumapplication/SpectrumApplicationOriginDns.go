// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spectrumapplication


type SpectrumApplicationOriginDns struct {
	// The name of the DNS record associated with the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/spectrum_application#name SpectrumApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The TTL of our resolution of your DNS record in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/spectrum_application#ttl SpectrumApplication#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// The type of DNS record associated with the origin.
	//
	// "" is used to specify a combination of A/AAAA records.
	// Available values: "", "A", "AAAA", "SRV".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/spectrum_application#type SpectrumApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

