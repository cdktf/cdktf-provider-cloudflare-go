// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package origincacertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OriginCaCertificateConfig struct {
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
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/origin_ca_certificate#csr OriginCaCertificate#csr}
	Csr *string `field:"optional" json:"csr" yaml:"csr"`
	// Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/origin_ca_certificate#hostnames OriginCaCertificate#hostnames}
	Hostnames *[]*string `field:"optional" json:"hostnames" yaml:"hostnames"`
	// The number of days for which the certificate should be valid. Available values: 7, 30, 90, 365, 730, 1095, 5475.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/origin_ca_certificate#requested_validity OriginCaCertificate#requested_validity}
	RequestedValidity *float64 `field:"optional" json:"requestedValidity" yaml:"requestedValidity"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa), or "keyless-certificate" (for Keyless SSL servers). Available values: "origin-rsa", "origin-ecc", "keyless-certificate".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/origin_ca_certificate#request_type OriginCaCertificate#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
}

