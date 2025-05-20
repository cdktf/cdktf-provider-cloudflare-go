// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigMtls struct {
	// Define CA certificate ID obtained after uploading CA cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/hyperdrive_config#ca_certificate_id HyperdriveConfig#ca_certificate_id}
	CaCertificateId *string `field:"optional" json:"caCertificateId" yaml:"caCertificateId"`
	// Define mTLS certificate ID obtained after uploading client cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/hyperdrive_config#mtls_certificate_id HyperdriveConfig#mtls_certificate_id}
	MtlsCertificateId *string `field:"optional" json:"mtlsCertificateId" yaml:"mtlsCertificateId"`
	// Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/hyperdrive_config#sslmode HyperdriveConfig#sslmode}
	Sslmode *string `field:"optional" json:"sslmode" yaml:"sslmode"`
}

