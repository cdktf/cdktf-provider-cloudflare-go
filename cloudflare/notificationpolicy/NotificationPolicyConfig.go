// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#account_id NotificationPolicy#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Refers to which event will trigger a Notification dispatch.
	//
	// You can use the endpoint to get available alert types which then will give you a list of possible values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#alert_type NotificationPolicy#alert_type}
	AlertType *string `field:"required" json:"alertType" yaml:"alertType"`
	// List of IDs that will be used when dispatching a notification.
	//
	// IDs for email type will be the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#mechanisms NotificationPolicy#mechanisms}
	Mechanisms *NotificationPolicyMechanisms `field:"required" json:"mechanisms" yaml:"mechanisms"`
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#name NotificationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional specification of how often to re-alert from the same incident, not support on all alert types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#alert_interval NotificationPolicy#alert_interval}
	AlertInterval *string `field:"optional" json:"alertInterval" yaml:"alertInterval"`
	// Optional description for the Notification policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#description NotificationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether or not the Notification policy is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#enabled NotificationPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria.
	//
	// This is only available for select alert types. See alert type documentation for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/notification_policy#filters NotificationPolicy#filters}
	Filters *NotificationPolicyFilters `field:"optional" json:"filters" yaml:"filters"`
}

