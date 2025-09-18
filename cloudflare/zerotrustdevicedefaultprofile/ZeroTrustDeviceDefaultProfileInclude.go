// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdevicedefaultprofile


type ZeroTrustDeviceDefaultProfileInclude struct {
	// The address in CIDR format to include in the tunnel. If `address` is present, `host` must not be present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_device_default_profile#address ZeroTrustDeviceDefaultProfile#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// A description of the Split Tunnel item, displayed in the client UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_device_default_profile#description ZeroTrustDeviceDefaultProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain name to include in the tunnel. If `host` is present, `address` must not be present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/zero_trust_device_default_profile#host ZeroTrustDeviceDefaultProfile#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
}

