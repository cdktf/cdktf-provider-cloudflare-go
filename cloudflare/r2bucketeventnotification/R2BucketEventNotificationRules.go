// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package r2bucketeventnotification


type R2BucketEventNotificationRules struct {
	// Array of R2 object actions that will trigger notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_event_notification#actions R2BucketEventNotification#actions}
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// A description that can be used to identify the event notification rule after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_event_notification#description R2BucketEventNotification#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Notifications will be sent only for objects with this prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_event_notification#prefix R2BucketEventNotification#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Notifications will be sent only for objects with this suffix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/r2_bucket_event_notification#suffix R2BucketEventNotification#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

