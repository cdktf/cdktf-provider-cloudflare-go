// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamliveinput

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamLiveInputConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#account_id StreamLiveInput#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Sets the creator ID asssociated with this live input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#default_creator StreamLiveInput#default_creator}
	DefaultCreator *string `field:"optional" json:"defaultCreator" yaml:"defaultCreator"`
	// Indicates the number of days after which the live inputs recordings will be deleted.
	//
	// When a stream completes and the recording is ready, the value is used to calculate a scheduled deletion date for that recording. Omit the field to indicate no change, or include with a `null` value to remove an existing scheduled deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#delete_recording_after_days StreamLiveInput#delete_recording_after_days}
	DeleteRecordingAfterDays *float64 `field:"optional" json:"deleteRecordingAfterDays" yaml:"deleteRecordingAfterDays"`
	// A unique identifier for a live input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#live_input_identifier StreamLiveInput#live_input_identifier}
	LiveInputIdentifier *string `field:"optional" json:"liveInputIdentifier" yaml:"liveInputIdentifier"`
	// A user modifiable key-value store used to reference other systems of record for managing live inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#meta StreamLiveInput#meta}
	Meta *string `field:"optional" json:"meta" yaml:"meta"`
	// Records the input to a Cloudflare Stream video.
	//
	// Behavior depends on the mode. In most cases, the video will initially be viewable as a live video and transition to on-demand after a condition is satisfied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/stream_live_input#recording StreamLiveInput#recording}
	Recording *StreamLiveInputRecording `field:"optional" json:"recording" yaml:"recording"`
}

