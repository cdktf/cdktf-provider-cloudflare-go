// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#enabled Worker#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#head_sampling_rate Worker#head_sampling_rate}
	HeadSamplingRate *float64 `field:"optional" json:"headSamplingRate" yaml:"headSamplingRate"`
	// Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#invocation_logs Worker#invocation_logs}
	InvocationLogs interface{} `field:"optional" json:"invocationLogs" yaml:"invocationLogs"`
}

