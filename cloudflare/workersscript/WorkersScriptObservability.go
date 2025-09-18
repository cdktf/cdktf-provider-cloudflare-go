// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptObservability struct {
	// Whether observability is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_script#enabled WorkersScript#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_script#head_sampling_rate WorkersScript#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Log settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_script#logs WorkersScript#logs}
	Logs *WorkersScriptObservabilityLogs `field:"optional" json:"logs" yaml:"logs"`
}

