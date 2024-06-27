// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accessmutualtlshostnamesettings


type AccessMutualTlsHostnameSettingsSettings struct {
	// The hostname that these settings apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#hostname AccessMutualTlsHostnameSettings#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Request client certificates for this hostname in China.
	//
	// Can only be set to true if this zone is china network enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#china_network AccessMutualTlsHostnameSettings#china_network}
	ChinaNetwork interface{} `field:"optional" json:"chinaNetwork" yaml:"chinaNetwork"`
	// Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.36.0/docs/resources/access_mutual_tls_hostname_settings#client_certificate_forwarding AccessMutualTlsHostnameSettings#client_certificate_forwarding}
	ClientCertificateForwarding interface{} `field:"optional" json:"clientCertificateForwarding" yaml:"clientCertificateForwarding"`
}

