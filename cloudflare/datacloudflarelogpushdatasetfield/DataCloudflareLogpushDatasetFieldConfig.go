// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarelogpushdatasetfield

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareLogpushDatasetFieldConfig struct {
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
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/logpush_dataset_field#account_id DataCloudflareLogpushDatasetField#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Name of the dataset.
	//
	// A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	// Available values: "access_requests", "audit_logs", "biso_user_actions", "casb_findings", "device_posture_results", "dlp_forensic_copies", "dns_firewall_logs", "dns_logs", "email_security_alerts", "firewall_events", "gateway_dns", "gateway_http", "gateway_network", "http_requests", "magic_ids_detections", "nel_reports", "network_analytics_logs", "page_shield_events", "sinkhole_http_logs", "spectrum_events", "ssh_logs", "workers_trace_events", "zaraz_events", "zero_trust_network_sessions".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/logpush_dataset_field#dataset_id DataCloudflareLogpushDatasetField#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/data-sources/logpush_dataset_field#zone_id DataCloudflareLogpushDatasetField#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

