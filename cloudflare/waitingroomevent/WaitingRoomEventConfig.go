// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroomevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaitingRoomEventConfig struct {
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
	// An ISO 8601 timestamp that marks the end of the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#event_end_time WaitingRoomEvent#event_end_time}
	EventEndTime *string `field:"required" json:"eventEndTime" yaml:"eventEndTime"`
	// An ISO 8601 timestamp that marks the start of the event.
	//
	// At this time, queued users will be processed with the event's configuration. The start time must be at least one minute before `event_end_time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#event_start_time WaitingRoomEvent#event_start_time}
	EventStartTime *string `field:"required" json:"eventStartTime" yaml:"eventStartTime"`
	// A unique name to identify the event. Only alphanumeric characters, hyphens and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#name WaitingRoomEvent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#waiting_room_id WaitingRoomEvent#waiting_room_id}.
	WaitingRoomId *string `field:"required" json:"waitingRoomId" yaml:"waitingRoomId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#zone_id WaitingRoomEvent#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// If set, the event will override the waiting room's `custom_page_html` property while it is active.
	//
	// If null, the event will inherit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#custom_page_html WaitingRoomEvent#custom_page_html}
	CustomPageHtml *string `field:"optional" json:"customPageHtml" yaml:"customPageHtml"`
	// A note that you can use to add more details about the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#description WaitingRoomEvent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, the event will override the waiting room's `disable_session_renewal` property while it is active.
	//
	// If null, the event will inherit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#disable_session_renewal WaitingRoomEvent#disable_session_renewal}
	DisableSessionRenewal interface{} `field:"optional" json:"disableSessionRenewal" yaml:"disableSessionRenewal"`
	// If set, the event will override the waiting room's `new_users_per_minute` property while it is active.
	//
	// If null, the event will inherit it. This can only be set if the event's `total_active_users` property is also set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#new_users_per_minute WaitingRoomEvent#new_users_per_minute}
	NewUsersPerMinute *float64 `field:"optional" json:"newUsersPerMinute" yaml:"newUsersPerMinute"`
	// An ISO 8601 timestamp that marks when to begin queueing all users before the event starts.
	//
	// The prequeue must start at least five minutes before `event_start_time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#prequeue_start_time WaitingRoomEvent#prequeue_start_time}
	PrequeueStartTime *string `field:"optional" json:"prequeueStartTime" yaml:"prequeueStartTime"`
	// If set, the event will override the waiting room's `queueing_method` property while it is active.
	//
	// If null, the event will inherit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#queueing_method WaitingRoomEvent#queueing_method}
	QueueingMethod *string `field:"optional" json:"queueingMethod" yaml:"queueingMethod"`
	// If set, the event will override the waiting room's `session_duration` property while it is active.
	//
	// If null, the event will inherit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#session_duration WaitingRoomEvent#session_duration}
	SessionDuration *float64 `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// If enabled, users in the prequeue will be shuffled randomly at the `event_start_time`.
	//
	// Requires that `prequeue_start_time` is not null. This is useful for situations when many users will join the event prequeue at the same time and you want to shuffle them to ensure fairness. Naturally, it makes the most sense to enable this feature when the `queueing_method` during the event respects ordering such as **fifo**, or else the shuffling may be unnecessary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#shuffle_at_event_start WaitingRoomEvent#shuffle_at_event_start}
	ShuffleAtEventStart interface{} `field:"optional" json:"shuffleAtEventStart" yaml:"shuffleAtEventStart"`
	// Suspends or allows an event.
	//
	// If set to `true`, the event is ignored and traffic will be handled based on the waiting room configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#suspended WaitingRoomEvent#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// If set, the event will override the waiting room's `total_active_users` property while it is active.
	//
	// If null, the event will inherit it. This can only be set if the event's `new_users_per_minute` property is also set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/waiting_room_event#total_active_users WaitingRoomEvent#total_active_users}
	TotalActiveUsers *float64 `field:"optional" json:"totalActiveUsers" yaml:"totalActiveUsers"`
}

