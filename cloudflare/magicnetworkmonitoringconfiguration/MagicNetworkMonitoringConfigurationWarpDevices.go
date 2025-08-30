// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicnetworkmonitoringconfiguration


type MagicNetworkMonitoringConfigurationWarpDevices struct {
	// Unique identifier for the warp device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_network_monitoring_configuration#id MagicNetworkMonitoringConfiguration#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Name of the warp device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_network_monitoring_configuration#name MagicNetworkMonitoringConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device. Only /32 addresses are currently supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/magic_network_monitoring_configuration#router_ip MagicNetworkMonitoringConfiguration#router_ip}
	RouterIp *string `field:"required" json:"routerIp" yaml:"routerIp"`
}

