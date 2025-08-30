// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionPlacement struct {
	// Placement mode for the version. Available values: "smart".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/worker_version#mode WorkerVersion#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

