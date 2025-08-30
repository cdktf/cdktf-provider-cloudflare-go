// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy


type NotificationPolicyFilters struct {
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#actions NotificationPolicy#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Used for configuring radar_notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#affected_asns NotificationPolicy#affected_asns}
	AffectedAsns *[]*string `field:"optional" json:"affectedAsns" yaml:"affectedAsns"`
	// Used for configuring incident_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#affected_components NotificationPolicy#affected_components}
	AffectedComponents *[]*string `field:"optional" json:"affectedComponents" yaml:"affectedComponents"`
	// Used for configuring radar_notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#affected_locations NotificationPolicy#affected_locations}
	AffectedLocations *[]*string `field:"optional" json:"affectedLocations" yaml:"affectedLocations"`
	// Used for configuring maintenance_event_notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#airport_code NotificationPolicy#airport_code}
	AirportCode *[]*string `field:"optional" json:"airportCode" yaml:"airportCode"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#alert_trigger_preferences NotificationPolicy#alert_trigger_preferences}
	AlertTriggerPreferences *[]*string `field:"optional" json:"alertTriggerPreferences" yaml:"alertTriggerPreferences"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#alert_trigger_preferences_value NotificationPolicy#alert_trigger_preferences_value}
	AlertTriggerPreferencesValue *[]*string `field:"optional" json:"alertTriggerPreferencesValue" yaml:"alertTriggerPreferencesValue"`
	// Used for configuring load_balancing_pool_enablement_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#enabled NotificationPolicy#enabled}
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// Used for configuring pages_event_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#environment NotificationPolicy#environment}
	Environment *[]*string `field:"optional" json:"environment" yaml:"environment"`
	// Used for configuring pages_event_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#event NotificationPolicy#event}
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// Used for configuring load_balancing_health_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#event_source NotificationPolicy#event_source}
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#event_type NotificationPolicy#event_type}
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#group_by NotificationPolicy#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Used for configuring health_check_status_notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#health_check_id NotificationPolicy#health_check_id}
	HealthCheckId *[]*string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Used for configuring incident_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#incident_impact NotificationPolicy#incident_impact}
	IncidentImpact *[]*string `field:"optional" json:"incidentImpact" yaml:"incidentImpact"`
	// Used for configuring stream_live_notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#input_id NotificationPolicy#input_id}
	InputId *[]*string `field:"optional" json:"inputId" yaml:"inputId"`
	// Used for configuring security_insights_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#insight_class NotificationPolicy#insight_class}
	InsightClass *[]*string `field:"optional" json:"insightClass" yaml:"insightClass"`
	// Used for configuring billing_usage_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#limit NotificationPolicy#limit}
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// Used for configuring logo_match_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#logo_tag NotificationPolicy#logo_tag}
	LogoTag *[]*string `field:"optional" json:"logoTag" yaml:"logoTag"`
	// Used for configuring advanced_ddos_attack_l4_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#megabits_per_second NotificationPolicy#megabits_per_second}
	MegabitsPerSecond *[]*string `field:"optional" json:"megabitsPerSecond" yaml:"megabitsPerSecond"`
	// Used for configuring load_balancing_health_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#new_health NotificationPolicy#new_health}
	NewHealth *[]*string `field:"optional" json:"newHealth" yaml:"newHealth"`
	// Used for configuring tunnel_health_event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#new_status NotificationPolicy#new_status}
	NewStatus *[]*string `field:"optional" json:"newStatus" yaml:"newStatus"`
	// Used for configuring advanced_ddos_attack_l4_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#packets_per_second NotificationPolicy#packets_per_second}
	PacketsPerSecond *[]*string `field:"optional" json:"packetsPerSecond" yaml:"packetsPerSecond"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#pool_id NotificationPolicy#pool_id}
	PoolId *[]*string `field:"optional" json:"poolId" yaml:"poolId"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#pop_names NotificationPolicy#pop_names}
	PopNames *[]*string `field:"optional" json:"popNames" yaml:"popNames"`
	// Used for configuring billing_usage_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#product NotificationPolicy#product}
	Product *[]*string `field:"optional" json:"product" yaml:"product"`
	// Used for configuring pages_event_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#project_id NotificationPolicy#project_id}
	ProjectId *[]*string `field:"optional" json:"projectId" yaml:"projectId"`
	// Used for configuring advanced_ddos_attack_l4_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#protocol NotificationPolicy#protocol}
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#query_tag NotificationPolicy#query_tag}
	QueryTag *[]*string `field:"optional" json:"queryTag" yaml:"queryTag"`
	// Used for configuring advanced_ddos_attack_l7_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#requests_per_second NotificationPolicy#requests_per_second}
	RequestsPerSecond *[]*string `field:"optional" json:"requestsPerSecond" yaml:"requestsPerSecond"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#selectors NotificationPolicy#selectors}
	Selectors *[]*string `field:"optional" json:"selectors" yaml:"selectors"`
	// Used for configuring clickhouse_alert_fw_ent_anomaly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#services NotificationPolicy#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#slo NotificationPolicy#slo}
	Slo *[]*string `field:"optional" json:"slo" yaml:"slo"`
	// Used for configuring health_check_status_notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#status NotificationPolicy#status}
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// Used for configuring advanced_ddos_attack_l7_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#target_hostname NotificationPolicy#target_hostname}
	TargetHostname *[]*string `field:"optional" json:"targetHostname" yaml:"targetHostname"`
	// Used for configuring advanced_ddos_attack_l4_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#target_ip NotificationPolicy#target_ip}
	TargetIp *[]*string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// Used for configuring advanced_ddos_attack_l7_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#target_zone_name NotificationPolicy#target_zone_name}
	TargetZoneName *[]*string `field:"optional" json:"targetZoneName" yaml:"targetZoneName"`
	// Used for configuring traffic_anomalies_alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#traffic_exclusions NotificationPolicy#traffic_exclusions}
	TrafficExclusions *[]*string `field:"optional" json:"trafficExclusions" yaml:"trafficExclusions"`
	// Used for configuring tunnel_health_event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#tunnel_id NotificationPolicy#tunnel_id}
	TunnelId *[]*string `field:"optional" json:"tunnelId" yaml:"tunnelId"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#tunnel_name NotificationPolicy#tunnel_name}
	TunnelName *[]*string `field:"optional" json:"tunnelName" yaml:"tunnelName"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#where NotificationPolicy#where}
	Where *[]*string `field:"optional" json:"where" yaml:"where"`
	// Usage depends on specific alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/notification_policy#zones NotificationPolicy#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

