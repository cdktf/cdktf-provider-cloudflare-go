// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigMtls struct {
	// CA certificate ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#ca_certificate_id HyperdriveConfig#ca_certificate_id}
	CaCertificateId *string `field:"optional" json:"caCertificateId" yaml:"caCertificateId"`
	// mTLS certificate ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#mtls_certificate_id HyperdriveConfig#mtls_certificate_id}
	MtlsCertificateId *string `field:"optional" json:"mtlsCertificateId" yaml:"mtlsCertificateId"`
	// SSL mode used for CA verification. Must be 'require', 'verify-ca', or 'verify-full'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/hyperdrive_config#sslmode HyperdriveConfig#sslmode}
	Sslmode *string `field:"optional" json:"sslmode" yaml:"sslmode"`
}

