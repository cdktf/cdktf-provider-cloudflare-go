// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keylesscertificate


type KeylessCertificateTunnel struct {
	// Private IP of the Key Server Host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/keyless_certificate#private_ip KeylessCertificate#private_ip}
	PrivateIp *string `field:"required" json:"privateIp" yaml:"privateIp"`
	// Cloudflare Tunnel Virtual Network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/keyless_certificate#vnet_id KeylessCertificate#vnet_id}
	VnetId *string `field:"required" json:"vnetId" yaml:"vnetId"`
}

