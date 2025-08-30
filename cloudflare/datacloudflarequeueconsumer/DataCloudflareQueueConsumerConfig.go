// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarequeueconsumer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareQueueConsumerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/queue_consumer#account_id DataCloudflareQueueConsumer#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// A Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/queue_consumer#consumer_id DataCloudflareQueueConsumer#consumer_id}
	ConsumerId *string `field:"required" json:"consumerId" yaml:"consumerId"`
	// A Resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/queue_consumer#queue_id DataCloudflareQueueConsumer#queue_id}
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

