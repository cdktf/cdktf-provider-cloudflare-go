// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerscript


type WorkerScriptPlacement struct {
	// The placement mode for the Worker. Available values: `smart`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/worker_script#mode WorkerScript#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

