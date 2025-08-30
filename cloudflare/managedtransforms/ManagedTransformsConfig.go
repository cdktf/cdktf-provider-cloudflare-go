// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedtransforms

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedTransformsConfig struct {
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
	// The list of Managed Request Transforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/managed_transforms#managed_request_headers ManagedTransforms#managed_request_headers}
	ManagedRequestHeaders interface{} `field:"required" json:"managedRequestHeaders" yaml:"managedRequestHeaders"`
	// The list of Managed Response Transforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/managed_transforms#managed_response_headers ManagedTransforms#managed_response_headers}
	ManagedResponseHeaders interface{} `field:"required" json:"managedResponseHeaders" yaml:"managedResponseHeaders"`
	// The unique ID of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/managed_transforms#zone_id ManagedTransforms#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

