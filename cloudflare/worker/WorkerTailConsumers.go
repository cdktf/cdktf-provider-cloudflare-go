// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worker


type WorkerTailConsumers struct {
	// Name of the consumer Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/worker#name Worker#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

