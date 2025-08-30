// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflarespectrumapplications

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareSpectrumApplicationsConfig struct {
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
	// Zone identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/spectrum_applications#zone_id DataCloudflareSpectrumApplications#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Sets the direction by which results are ordered. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/spectrum_applications#direction DataCloudflareSpectrumApplications#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/spectrum_applications#max_items DataCloudflareSpectrumApplications#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Application field by which results are ordered. Available values: "protocol", "app_id", "created_on", "modified_on", "dns".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/spectrum_applications#order DataCloudflareSpectrumApplications#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

