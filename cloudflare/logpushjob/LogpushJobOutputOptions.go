// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logpushjob


type LogpushJobOutputOptions struct {
	// String to be prepended before each batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#batch_prefix LogpushJob#batch_prefix}
	BatchPrefix *string `field:"optional" json:"batchPrefix" yaml:"batchPrefix"`
	// String to be appended after each batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#batch_suffix LogpushJob#batch_suffix}
	BatchSuffix *string `field:"optional" json:"batchSuffix" yaml:"batchSuffix"`
	// Mitigation for CVE-2021-44228.
	//
	// If set to true, will cause all occurrences of ${ in the generated files to be replaced with x{. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#cve20214428 LogpushJob#cve20214428}
	Cve20214428 interface{} `field:"optional" json:"cve20214428" yaml:"cve20214428"`
	// String to join fields. This field be ignored when record_template is set. Defaults to `,`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#field_delimiter LogpushJob#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// List of field names to be included in the Logpush output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#field_names LogpushJob#field_names}
	FieldNames *[]*string `field:"optional" json:"fieldNames" yaml:"fieldNames"`
	// Specifies the output type. Available values: `ndjson`, `csv`. Defaults to `ndjson`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#output_type LogpushJob#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// String to be inserted in-between the records as separator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#record_delimiter LogpushJob#record_delimiter}
	RecordDelimiter *string `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
	// String to be prepended before each record. Defaults to `{`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#record_prefix LogpushJob#record_prefix}
	RecordPrefix *string `field:"optional" json:"recordPrefix" yaml:"recordPrefix"`
	// String to be appended after each record. Defaults to `} `.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#record_suffix LogpushJob#record_suffix}
	RecordSuffix *string `field:"optional" json:"recordSuffix" yaml:"recordSuffix"`
	// String to use as template for each record instead of the default comma-separated list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#record_template LogpushJob#record_template}
	RecordTemplate *string `field:"optional" json:"recordTemplate" yaml:"recordTemplate"`
	// Specifies the sampling rate. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#sample_rate LogpushJob#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// Specifies the format for timestamps. Available values: `unixnano`, `unix`, `rfc3339`. Defaults to `unixnano`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/logpush_job#timestamp_format LogpushJob#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

