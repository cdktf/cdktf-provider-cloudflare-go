// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamwatermark

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamWatermarkAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#account_id StreamWatermarkA#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The image file to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#file StreamWatermarkA#file}
	File *string `field:"required" json:"file" yaml:"file"`
	// The unique identifier for a watermark profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#identifier StreamWatermarkA#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// A short description of the watermark profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#name StreamWatermarkA#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The translucency of the image.
	//
	// A value of `0.0` makes the image completely transparent, and `1.0` makes the image completely opaque. Note that if the image is already semi-transparent, setting this to `1.0` will not make the image completely opaque.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#opacity StreamWatermarkA#opacity}
	Opacity *float64 `field:"optional" json:"opacity" yaml:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video and the image.
	//
	// `0.0` indicates no padding, and `1.0` indicates a fully padded video width or length, as determined by the algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#padding StreamWatermarkA#padding}
	Padding *float64 `field:"optional" json:"padding" yaml:"padding"`
	// The location of the image.
	//
	// Valid positions are: `upperRight`, `upperLeft`, `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the `padding` parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#position StreamWatermarkA#position}
	Position *string `field:"optional" json:"position" yaml:"position"`
	// The size of the image relative to the overall size of the video.
	//
	// This parameter will adapt to horizontal and vertical videos automatically. `0.0` indicates no scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/stream_watermark#scale StreamWatermarkA#scale}
	Scale *float64 `field:"optional" json:"scale" yaml:"scale"`
}

