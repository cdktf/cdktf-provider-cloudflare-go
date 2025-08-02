// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logpushjob


type LogpushJobOutputOptions struct {
	// String to be prepended before each batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#batch_prefix LogpushJob#batch_prefix}
	BatchPrefix *string `field:"optional" json:"batchPrefix" yaml:"batchPrefix"`
	// String to be appended after each batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#batch_suffix LogpushJob#batch_suffix}
	BatchSuffix *string `field:"optional" json:"batchSuffix" yaml:"batchSuffix"`
	// If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#cve_2021_44228 LogpushJob#cve_2021_44228}
	Cve202144228 interface{} `field:"optional" json:"cve202144228" yaml:"cve202144228"`
	// String to join fields. This field be ignored when `record_template` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#field_delimiter LogpushJob#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// List of field names to be included in the Logpush output.
	//
	// For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#field_names LogpushJob#field_names}
	FieldNames *[]*string `field:"optional" json:"fieldNames" yaml:"fieldNames"`
	// Specifies the output type, such as `ndjson` or `csv`.
	//
	// This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	// Available values: "ndjson", "csv".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#output_type LogpushJob#output_type}
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// String to be inserted in-between the records as separator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#record_delimiter LogpushJob#record_delimiter}
	RecordDelimiter *string `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
	// String to be prepended before each record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#record_prefix LogpushJob#record_prefix}
	RecordPrefix *string `field:"optional" json:"recordPrefix" yaml:"recordPrefix"`
	// String to be appended after each record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#record_suffix LogpushJob#record_suffix}
	RecordSuffix *string `field:"optional" json:"recordSuffix" yaml:"recordSuffix"`
	// String to use as template for each record instead of the default json key value mapping.
	//
	// All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#record_template LogpushJob#record_template}
	RecordTemplate *string `field:"optional" json:"recordTemplate" yaml:"recordTemplate"`
	// Floating number to specify sampling rate.
	//
	// Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#sample_rate LogpushJob#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`. Available values: "unixnano", "unix", "rfc3339".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/logpush_job#timestamp_format LogpushJob#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

