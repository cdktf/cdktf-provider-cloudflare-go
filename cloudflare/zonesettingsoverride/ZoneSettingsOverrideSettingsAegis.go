// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zonesettingsoverride


type ZoneSettingsOverrideSettingsAegis struct {
	// Whether Aegis zone setting is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zone_settings_override#enabled ZoneSettingsOverride#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which Cloudflare will connect to origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zone_settings_override#pool_id ZoneSettingsOverride#pool_id}
	PoolId *string `field:"optional" json:"poolId" yaml:"poolId"`
}

