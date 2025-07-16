// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicywebhooks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyWebhooksConfig struct {
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
	// The account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/notification_policy_webhooks#account_id NotificationPolicyWebhooks#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the webhook destination.
	//
	// This will be included in the request body when you receive a webhook notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/notification_policy_webhooks#name NotificationPolicyWebhooks#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The POST endpoint to call when dispatching a notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/notification_policy_webhooks#url NotificationPolicyWebhooks#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when dispatching generic webhook notifications or formatted for supported destinations.
	//
	// Secrets are not returned in any API response body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/notification_policy_webhooks#secret NotificationPolicyWebhooks#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

