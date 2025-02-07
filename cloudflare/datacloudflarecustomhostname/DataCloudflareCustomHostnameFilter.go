// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarecustomhostname


type DataCloudflareCustomHostnameFilter struct {
	// Direction to order hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/custom_hostname#direction DataCloudflareCustomHostname#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with the 'id' parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/custom_hostname#hostname DataCloudflareCustomHostname#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Hostname ID to match against.
	//
	// This ID was generated and returned during the initial custom_hostname creation. This parameter cannot be used with the 'hostname' parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/custom_hostname#id DataCloudflareCustomHostname#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Field to order hostnames by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/custom_hostname#order DataCloudflareCustomHostname#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Whether to filter hostnames based on if they have SSL enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/data-sources/custom_hostname#ssl DataCloudflareCustomHostname#ssl}
	Ssl *float64 `field:"optional" json:"ssl" yaml:"ssl"`
}

