// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customhostname


type CustomHostnameSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/custom_hostname#ciphers CustomHostname#ciphers}
	Ciphers *[]*string `field:"optional" json:"ciphers" yaml:"ciphers"`
	// Whether or not Early Hints is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/custom_hostname#early_hints CustomHostname#early_hints}
	EarlyHints *string `field:"optional" json:"earlyHints" yaml:"earlyHints"`
	// Whether or not HTTP2 is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/custom_hostname#http2 CustomHostname#http2}
	Http2 *string `field:"optional" json:"http2" yaml:"http2"`
	// The minimum TLS version supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/custom_hostname#min_tls_version CustomHostname#min_tls_version}
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Whether or not TLS 1.3 is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/custom_hostname#tls_1_3 CustomHostname#tls_1_3}
	Tls13 *string `field:"optional" json:"tls13" yaml:"tls13"`
}

