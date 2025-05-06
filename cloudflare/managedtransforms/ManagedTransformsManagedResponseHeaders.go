// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedtransforms


type ManagedTransformsManagedResponseHeaders struct {
	// Whether the Managed Transform is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/managed_transforms#enabled ManagedTransforms#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The human-readable identifier of the Managed Transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/managed_transforms#id ManagedTransforms#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

