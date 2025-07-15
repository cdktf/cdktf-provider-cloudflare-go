// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyConfig struct {
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
	// The account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#account_id NotificationPolicy#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Refers to which event will trigger a Notification dispatch.
	//
	// You can use the endpoint to get available alert types which then will give you a list of possible values.
	// Available values: "access_custom_certificate_expiration_type", "advanced_ddos_attack_l4_alert", "advanced_ddos_attack_l7_alert", "advanced_http_alert_error", "bgp_hijack_notification", "billing_usage_alert", "block_notification_block_removed", "block_notification_new_block", "block_notification_review_rejected", "bot_traffic_basic_alert", "brand_protection_alert", "brand_protection_digest", "clickhouse_alert_fw_anomaly", "clickhouse_alert_fw_ent_anomaly", "cloudforce_one_request_notification", "custom_analytics", "custom_bot_detection_alert", "custom_ssl_certificate_event_type", "dedicated_ssl_certificate_event_type", "device_connectivity_anomaly_alert", "dos_attack_l4", "dos_attack_l7", "expiring_service_token_alert", "failing_logpush_job_disabled_alert", "fbm_auto_advertisement", "fbm_dosd_attack", "fbm_volumetric_attack", "health_check_status_notification", "hostname_aop_custom_certificate_expiration_type", "http_alert_edge_error", "http_alert_origin_error", "image_notification", "image_resizing_notification", "incident_alert", "load_balancing_health_alert", "load_balancing_pool_enablement_alert", "logo_match_alert", "magic_tunnel_health_check_event", "magic_wan_tunnel_health", "maintenance_event_notification", "mtls_certificate_store_certificate_expiration_type", "pages_event_alert", "radar_notification", "real_origin_monitoring", "scriptmonitor_alert_new_code_change_detections", "scriptmonitor_alert_new_hosts", "scriptmonitor_alert_new_malicious_hosts", "scriptmonitor_alert_new_malicious_scripts", "scriptmonitor_alert_new_malicious_url", "scriptmonitor_alert_new_max_length_resource_url", "scriptmonitor_alert_new_resources", "secondary_dns_all_primaries_failing", "secondary_dns_primaries_failing", "secondary_dns_warning", "secondary_dns_zone_successfully_updated", "secondary_dns_zone_validation_warning", "security_insights_alert", "sentinel_alert", "stream_live_notifications", "synthetic_test_latency_alert", "synthetic_test_low_availability_alert", "traffic_anomalies_alert", "tunnel_health_event", "tunnel_update_event", "universal_ssl_event_type", "web_analytics_metrics_update", "zone_aop_custom_certificate_expiration_type".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#alert_type NotificationPolicy#alert_type}
	AlertType *string `field:"required" json:"alertType" yaml:"alertType"`
	// List of IDs that will be used when dispatching a notification.
	//
	// IDs for email type will be the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#mechanisms NotificationPolicy#mechanisms}
	Mechanisms *NotificationPolicyMechanisms `field:"required" json:"mechanisms" yaml:"mechanisms"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#name NotificationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional specification of how often to re-alert from the same incident, not support on all alert types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#alert_interval NotificationPolicy#alert_interval}
	AlertInterval *string `field:"optional" json:"alertInterval" yaml:"alertInterval"`
	// Optional description for the Notification policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#description NotificationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether or not the Notification policy is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#enabled NotificationPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria.
	//
	// This is only available for select alert types. See alert type documentation for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/notification_policy#filters NotificationPolicy#filters}
	Filters *NotificationPolicyFilters `field:"optional" json:"filters" yaml:"filters"`
}

