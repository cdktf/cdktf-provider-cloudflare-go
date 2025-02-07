// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package queueconsumer


type QueueConsumerSettings struct {
	// The maximum number of messages to include in a batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#batch_size QueueConsumer#batch_size}
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Maximum number of concurrent consumers that may consume from this Queue.
	//
	// Set to `null` to automatically opt in to the platform's maximum (recommended).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#max_concurrency QueueConsumer#max_concurrency}
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#max_retries QueueConsumer#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The number of milliseconds to wait for a batch to fill up before attempting to deliver it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#max_wait_time_ms QueueConsumer#max_wait_time_ms}
	MaxWaitTimeMs *float64 `field:"optional" json:"maxWaitTimeMs" yaml:"maxWaitTimeMs"`
	// The number of seconds to delay before making the message available for another attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#retry_delay QueueConsumer#retry_delay}
	RetryDelay *float64 `field:"optional" json:"retryDelay" yaml:"retryDelay"`
	// The number of milliseconds that a message is exclusively leased.
	//
	// After the timeout, the message becomes available for another attempt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/queue_consumer#visibility_timeout_ms QueueConsumer#visibility_timeout_ms}
	VisibilityTimeoutMs *float64 `field:"optional" json:"visibilityTimeoutMs" yaml:"visibilityTimeoutMs"`
}

