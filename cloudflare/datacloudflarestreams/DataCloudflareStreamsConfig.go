// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarestreams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareStreamsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#account_id DataCloudflareStreams#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Lists videos in ascending order of creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#asc DataCloudflareStreams#asc}
	Asc interface{} `field:"optional" json:"asc" yaml:"asc"`
	// A user-defined identifier for the media creator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#creator DataCloudflareStreams#creator}
	Creator *string `field:"optional" json:"creator" yaml:"creator"`
	// Lists videos created before the specified date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#end DataCloudflareStreams#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// Includes the total number of videos associated with the submitted query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#include_counts DataCloudflareStreams#include_counts}
	IncludeCounts interface{} `field:"optional" json:"includeCounts" yaml:"includeCounts"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#max_items DataCloudflareStreams#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Searches over the `name` key in the `meta` field.
	//
	// This field can be set with or after the upload request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#search DataCloudflareStreams#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
	// Lists videos created after the specified date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#start DataCloudflareStreams#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
	// Specifies the processing status for all quality levels for a video. Available values: "pendingupload", "downloading", "queued", "inprogress", "ready", "error".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#status DataCloudflareStreams#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Specifies whether the video is `vod` or `live`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/streams#type DataCloudflareStreams#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

