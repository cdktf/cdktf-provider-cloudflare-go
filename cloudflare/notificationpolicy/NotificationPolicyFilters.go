// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy


type NotificationPolicyFilters struct {
	// Targeted actions for alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#actions NotificationPolicy#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Affected components for alert.
	//
	// Available values: `API`, `API Shield`, `Access`, `Always Online`, `Analytics`, `Apps Marketplace`, `Argo Smart Routing`, `Audit Logs`, `Authoritative DNS`, `Billing`, `Bot Management`, `Bring Your Own IP (BYOIP)`, `Browser Isolation`, `CDN Cache Purge`, `CDN/Cache`, `Cache Reserve`, `Challenge Platform`, `Cloud Access Security Broker (CASB)`, `Community Site`, `D1`, `DNS Root Servers`, `DNS Updates`, `Dashboard`, `Data Loss Prevention (DLP)`, `Developer's Site`, `Digital Experience Monitoring (DEX)`, `Distributed Web Gateway`, `Durable Objects`, `Email Routing`, `Ethereum Gateway`, `Firewall`, `Gateway`, `Geo-Key Manager`, `Image Resizing`, `Images`, `Infrastructure`, `Lists`, `Load Balancing and Monitoring`, `Logs`, `Magic Firewall`, `Magic Transit`, `Magic WAN`, `Magic WAN Connector`, `Marketing Site`, `Mirage`, `Network`, `Notifications`, `Observatory`, `Page Shield`, `Pages`, `R2`, `Radar`, `Randomness Beacon`, `Recursive DNS`, `Registrar`, `Registration Data Access Protocol (RDAP)`, `SSL Certificate Provisioning`, `SSL for SaaS Provisioning`, `Security Center`, `Snippets`, `Spectrum`, `Speed Optimizations`, `Stream`, `Support Site`, `Time Services`, `Trace`, `Tunnel`, `Turnstile`, `WARP`, `Waiting Room`, `Web Analytics`, `Workers`, `Workers KV`, `Workers Preview`, `Zaraz`, `Zero Trust`, `Zero Trust Dashboard`, `Zone Versioning`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#affected_components NotificationPolicy#affected_components}
	AffectedComponents *[]*string `field:"optional" json:"affectedComponents" yaml:"affectedComponents"`
	// Filter on Points of Presence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#airport_code NotificationPolicy#airport_code}
	AirportCode *[]*string `field:"optional" json:"airportCode" yaml:"airportCode"`
	// Alert trigger preferences. Example: `slo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#alert_trigger_preferences NotificationPolicy#alert_trigger_preferences}
	AlertTriggerPreferences *[]*string `field:"optional" json:"alertTriggerPreferences" yaml:"alertTriggerPreferences"`
	// State of the pool to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#enabled NotificationPolicy#enabled}
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// Environment of pages. Available values: `ENVIRONMENT_PREVIEW`, `ENVIRONMENT_PRODUCTION`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#environment NotificationPolicy#environment}
	Environment *[]*string `field:"optional" json:"environment" yaml:"environment"`
	// Pages event to alert. Available values: `EVENT_DEPLOYMENT_STARTED`, `EVENT_DEPLOYMENT_FAILED`, `EVENT_DEPLOYMENT_SUCCESS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#event NotificationPolicy#event}
	Event *[]*string `field:"optional" json:"event" yaml:"event"`
	// Source configuration to alert on for pool or origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#event_source NotificationPolicy#event_source}
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Stream event type to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#event_type NotificationPolicy#event_type}
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// Alert grouping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#group_by NotificationPolicy#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Identifier health check. Required when using `filters.0.status`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#health_check_id NotificationPolicy#health_check_id}
	HealthCheckId *[]*string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// The incident impact level that will trigger the dispatch of a notification. Available values: `INCIDENT_IMPACT_NONE`, `INCIDENT_IMPACT_MINOR`, `INCIDENT_IMPACT_MAJOR`, `INCIDENT_IMPACT_CRITICAL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#incident_impact NotificationPolicy#incident_impact}
	IncidentImpact *[]*string `field:"optional" json:"incidentImpact" yaml:"incidentImpact"`
	// Stream input id to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#input_id NotificationPolicy#input_id}
	InputId *[]*string `field:"optional" json:"inputId" yaml:"inputId"`
	// A numerical limit. Example: `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#limit NotificationPolicy#limit}
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// Megabits per second threshold for dos alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#megabits_per_second NotificationPolicy#megabits_per_second}
	MegabitsPerSecond *[]*string `field:"optional" json:"megabitsPerSecond" yaml:"megabitsPerSecond"`
	// Health status to alert on for pool or origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#new_health NotificationPolicy#new_health}
	NewHealth *[]*string `field:"optional" json:"newHealth" yaml:"newHealth"`
	// Tunnel health status to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#new_status NotificationPolicy#new_status}
	NewStatus *[]*string `field:"optional" json:"newStatus" yaml:"newStatus"`
	// Packets per second threshold for dos alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#packets_per_second NotificationPolicy#packets_per_second}
	PacketsPerSecond *[]*string `field:"optional" json:"packetsPerSecond" yaml:"packetsPerSecond"`
	// Load balancer pool identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#pool_id NotificationPolicy#pool_id}
	PoolId *[]*string `field:"optional" json:"poolId" yaml:"poolId"`
	// Product name. Available values: `worker_requests`, `worker_durable_objects_requests`, `worker_durable_objects_duration`, `worker_durable_objects_data_transfer`, `worker_durable_objects_stored_data`, `worker_durable_objects_storage_deletes`, `worker_durable_objects_storage_writes`, `worker_durable_objects_storage_reads`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#product NotificationPolicy#product}
	Product *[]*string `field:"optional" json:"product" yaml:"product"`
	// Identifier of pages project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#project_id NotificationPolicy#project_id}
	ProjectId *[]*string `field:"optional" json:"projectId" yaml:"projectId"`
	// Protocol to alert on for dos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#protocol NotificationPolicy#protocol}
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
	// Requests per second threshold for dos alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#requests_per_second NotificationPolicy#requests_per_second}
	RequestsPerSecond *[]*string `field:"optional" json:"requestsPerSecond" yaml:"requestsPerSecond"`
	// Selectors for alert. Valid options depend on the alert type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#selectors NotificationPolicy#selectors}
	Selectors *[]*string `field:"optional" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#services NotificationPolicy#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// A numerical limit. Example: `99.9`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#slo NotificationPolicy#slo}
	Slo *[]*string `field:"optional" json:"slo" yaml:"slo"`
	// Status to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#status NotificationPolicy#status}
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// Target host to alert on for dos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#target_hostname NotificationPolicy#target_hostname}
	TargetHostname *[]*string `field:"optional" json:"targetHostname" yaml:"targetHostname"`
	// Target ip to alert on for dos in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#target_ip NotificationPolicy#target_ip}
	TargetIp *[]*string `field:"optional" json:"targetIp" yaml:"targetIp"`
	// Target domain to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#target_zone_name NotificationPolicy#target_zone_name}
	TargetZoneName *[]*string `field:"optional" json:"targetZoneName" yaml:"targetZoneName"`
	// Tunnel IDs to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#tunnel_id NotificationPolicy#tunnel_id}
	TunnelId *[]*string `field:"optional" json:"tunnelId" yaml:"tunnelId"`
	// Tunnel Names to alert on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#tunnel_name NotificationPolicy#tunnel_name}
	TunnelName *[]*string `field:"optional" json:"tunnelName" yaml:"tunnelName"`
	// Filter for alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#where NotificationPolicy#where}
	Where *[]*string `field:"optional" json:"where" yaml:"where"`
	// A list of zone identifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/notification_policy#zones NotificationPolicy#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

