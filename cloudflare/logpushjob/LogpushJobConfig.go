// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logpushjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogpushJobConfig struct {
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
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	//
	// Additional configuration parameters supported by the destination may be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#destination_conf LogpushJob#destination_conf}
	DestinationConf *string `field:"required" json:"destinationConf" yaml:"destinationConf"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#account_id LogpushJob#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#dataset LogpushJob#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// Flag that indicates if the job is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#enabled LogpushJob#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// This field is deprecated.
	//
	// Please use `max_upload_*` parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#frequency LogpushJob#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge Log Delivery jobs.
	//
	// Currently, Edge Log Delivery is only supported for the `http_requests` dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#kind LogpushJob#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// This field is deprecated.
	//
	// Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#logpull_options LogpushJob#logpull_options}
	LogpullOptions *string `field:"optional" json:"logpullOptions" yaml:"logpullOptions"`
	// The maximum uncompressed file size of a batch of logs.
	//
	// This setting value must be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with `edge` as its kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#max_upload_bytes LogpushJob#max_upload_bytes}
	MaxUploadBytes *float64 `field:"optional" json:"maxUploadBytes" yaml:"maxUploadBytes"`
	// The maximum interval in seconds for log batches.
	//
	// This setting must be between 30 and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with `edge` as its kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#max_upload_interval_seconds LogpushJob#max_upload_interval_seconds}
	MaxUploadIntervalSeconds *float64 `field:"optional" json:"maxUploadIntervalSeconds" yaml:"maxUploadIntervalSeconds"`
	// The maximum number of log lines per batch.
	//
	// This setting must be between 1000 and 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with `edge` as its kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#max_upload_records LogpushJob#max_upload_records}
	MaxUploadRecords *float64 `field:"optional" json:"maxUploadRecords" yaml:"maxUploadRecords"`
	// Optional human readable job name.
	//
	// Not unique. Cloudflare suggests that you set this to a meaningful string, like the domain name, to make it easier to identify your job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#name LogpushJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The structured replacement for `logpull_options`. When including this field, the `logpull_option` field will be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#output_options LogpushJob#output_options}
	OutputOptions *LogpushJobOutputOptions `field:"optional" json:"outputOptions" yaml:"outputOptions"`
	// Ownership challenge token to prove destination ownership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#ownership_challenge LogpushJob#ownership_challenge}
	OwnershipChallenge *string `field:"optional" json:"ownershipChallenge" yaml:"ownershipChallenge"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/logpush_job#zone_id LogpushJob#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

