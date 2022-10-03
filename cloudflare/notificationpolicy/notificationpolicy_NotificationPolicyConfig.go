package notificationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#account_id NotificationPolicy#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The event type that will trigger the dispatch of a notification.
	//
	// See the developer documentation for descriptions of [available alert types](https://developers.cloudflare.com/fundamentals/notifications/notification-available/). Available values: `billing_usage_alert`, `health_check_status_notification`, `g6_pool_toggle_alert`, `real_origin_monitoring`, `universal_ssl_event_type`, `dedicated_ssl_certificate_event_type`, `custom_ssl_certificate_event_type`, `access_custom_certificate_expiration_type`, `zone_aop_custom_certificate_expiration_type`, `bgp_hijack_notification`, `http_alert_origin_error`, `workers_alert`, `weekly_account_overview`, `expiring_service_token_alert`, `secondary_dns_all_primaries_failing`, `secondary_dns_zone_validation_warning`, `secondary_dns_primaries_failing`, `secondary_dns_zone_successfully_updated`, `dos_attack_l7`, `dos_attack_l4`, `advanced_ddos_attack_l7_alert`, `advanced_ddos_attack_l4_alert`, `fbm_volumetric_attack`, `fbm_auto_advertisement`, `load_balancing_pool_enablement_alert`, `load_balancing_health_alert`, `g6_health_alert`, `http_alert_edge_error`, `clickhouse_alert_fw_anomaly`, `clickhouse_alert_fw_ent_anomaly`, `failing_logpush_job_disabled_alert`, `scriptmonitor_alert_new_hosts`, `scriptmonitor_alert_new_scripts`, `scriptmonitor_alert_new_malicious_scripts`, `scriptmonitor_alert_new_malicious_url`, `scriptmonitor_alert_new_code_change_detections`, `scriptmonitor_alert_new_max_length_script_url`, `scriptmonitor_alert_new_malicious_hosts`, `sentinel_alert`, `hostname_aop_custom_certificate_expiration_type`, `stream_live_notifications`, `block_notification_new_block`, `block_notification_review_rejected`, `block_notification_review_accepted`, `web_analytics_metrics_update`, `workers_uptime`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#alert_type NotificationPolicy#alert_type}
	AlertType *string `field:"required" json:"alertType" yaml:"alertType"`
	// The status of the notification policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#enabled NotificationPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the notification policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#name NotificationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the notification policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#description NotificationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#email_integration NotificationPolicy#email_integration}
	EmailIntegration interface{} `field:"optional" json:"emailIntegration" yaml:"emailIntegration"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#filters NotificationPolicy#filters}
	Filters *NotificationPolicyFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#id NotificationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// pagerduty_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#pagerduty_integration NotificationPolicy#pagerduty_integration}
	PagerdutyIntegration interface{} `field:"optional" json:"pagerdutyIntegration" yaml:"pagerdutyIntegration"`
	// webhooks_integration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#webhooks_integration NotificationPolicy#webhooks_integration}
	WebhooksIntegration interface{} `field:"optional" json:"webhooksIntegration" yaml:"webhooksIntegration"`
}

