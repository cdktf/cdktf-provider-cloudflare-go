// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type NotificationPolicyFilters struct {
	// State of the pool to alert on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#enabled NotificationPolicy#enabled}
	Enabled *[]*string `field:"optional" json:"enabled" yaml:"enabled"`
	// Identifier health check.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#health_check_id NotificationPolicy#health_check_id}
	HealthCheckId *[]*string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// A numerical limit. Example: `100`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#limit NotificationPolicy#limit}
	Limit *[]*string `field:"optional" json:"limit" yaml:"limit"`
	// Load balancer pool identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#pool_id NotificationPolicy#pool_id}
	PoolId *[]*string `field:"optional" json:"poolId" yaml:"poolId"`
	// Product name. Available values: `worker_requests`, `worker_durable_objects_requests`, `worker_durable_objects_duration`, `worker_durable_objects_data_transfer`, `worker_durable_objects_stored_data`, `worker_durable_objects_storage_deletes`, `worker_durable_objects_storage_writes`, `worker_durable_objects_storage_reads`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#product NotificationPolicy#product}
	Product *[]*string `field:"optional" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#services NotificationPolicy#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
	// A numerical limit. Example: `99.9`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#slo NotificationPolicy#slo}
	Slo *[]*string `field:"optional" json:"slo" yaml:"slo"`
	// Status to alert on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#status NotificationPolicy#status}
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// A list of zone identifiers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/notification_policy#zones NotificationPolicy#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

