// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkerscustomdomain


type DataCloudflareWorkersCustomDomainFilter struct {
	// Worker environment associated with the zone and hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/workers_custom_domain#environment DataCloudflareWorkersCustomDomain#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Hostname of the Worker Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/workers_custom_domain#hostname DataCloudflareWorkersCustomDomain#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Worker service associated with the zone and hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/workers_custom_domain#service DataCloudflareWorkersCustomDomain#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Identifier of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/workers_custom_domain#zone_id DataCloudflareWorkersCustomDomain#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
	// Name of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/data-sources/workers_custom_domain#zone_name DataCloudflareWorkersCustomDomain#zone_name}
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

