// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationpolicy


type NotificationPolicyMechanisms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/notification_policy#email NotificationPolicy#email}.
	Email interface{} `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/notification_policy#pagerduty NotificationPolicy#pagerduty}.
	Pagerduty interface{} `field:"optional" json:"pagerduty" yaml:"pagerduty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/notification_policy#webhooks NotificationPolicy#webhooks}.
	Webhooks interface{} `field:"optional" json:"webhooks" yaml:"webhooks"`
}

