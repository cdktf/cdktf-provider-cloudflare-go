// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamConfig struct {
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
	// The account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#account_id Stream#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Lists the origins allowed to display the video.
	//
	// Enter allowed origin domains in an array and use `*` for wildcard subdomains. Empty arrays allow the video to be viewed on any origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#allowed_origins Stream#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#creator Stream#creator}
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// A Cloudflare-generated unique identifier for a media item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#identifier Stream#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The maximum duration in seconds for a video upload.
	//
	// Can be set for a video that is not yet uploaded to limit its duration. Uploads that exceed the specified duration will fail during processing. A value of `-1` means the value is unknown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#max_duration_seconds Stream#max_duration_seconds}
	MaxDurationSeconds *float64 `field:"optional" json:"maxDurationSeconds" yaml:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for managing videos.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#meta Stream#meta}
	Meta *string `field:"optional" json:"meta" yaml:"meta"`
	// Indicates whether the video can be a accessed using the UID.
	//
	// When set to `true`, a signed token must be generated with a signing key to view the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#require_signed_urls Stream#require_signed_urls}
	RequireSignedUrls interface{} `field:"optional" json:"requireSignedUrls" yaml:"requireSignedUrls"`
	// Indicates the date and time at which the video will be deleted.
	//
	// Omit the field to indicate no change, or include with a `null` value to remove an existing scheduled deletion. If specified, must be at least 30 days from upload time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#scheduled_deletion Stream#scheduled_deletion}
	ScheduledDeletion *string `field:"optional" json:"scheduledDeletion" yaml:"scheduledDeletion"`
	// The timestamp for a thumbnail image calculated as a percentage value of the video's duration.
	//
	// To convert from a second-wise timestamp to a percentage, divide the desired timestamp by the total duration of the video.  If this value is not set, the default thumbnail image is taken from 0s of the video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#thumbnail_timestamp_pct Stream#thumbnail_timestamp_pct}
	ThumbnailTimestampPct *float64 `field:"optional" json:"thumbnailTimestampPct" yaml:"thumbnailTimestampPct"`
	// The date and time when the video upload URL is no longer valid for direct user uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/stream#upload_expiry Stream#upload_expiry}
	UploadExpiry *string `field:"optional" json:"uploadExpiry" yaml:"uploadExpiry"`
}

