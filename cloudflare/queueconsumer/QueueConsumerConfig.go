// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package queueconsumer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QueueConsumerConfig struct {
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
	// A Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#account_id QueueConsumer#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#queue_id QueueConsumer#queue_id}
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// A Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#consumer_id QueueConsumer#consumer_id}
	ConsumerId *string `field:"optional" json:"consumerId" yaml:"consumerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#dead_letter_queue QueueConsumer#dead_letter_queue}.
	DeadLetterQueue *string `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Name of a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#script_name QueueConsumer#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#settings QueueConsumer#settings}.
	Settings *QueueConsumerSettings `field:"optional" json:"settings" yaml:"settings"`
	// Available values: "worker", "http_pull".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/queue_consumer#type QueueConsumer#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

