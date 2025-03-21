// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package queue


type QueueSettings struct {
	// Number of seconds to delay delivery of all messages to consumers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/queue#delivery_delay Queue#delivery_delay}
	DeliveryDelay *float64 `field:"optional" json:"deliveryDelay" yaml:"deliveryDelay"`
	// Number of seconds after which an unconsumed message will be delayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/queue#message_retention_period Queue#message_retention_period}
	MessageRetentionPeriod *float64 `field:"optional" json:"messageRetentionPeriod" yaml:"messageRetentionPeriod"`
}

