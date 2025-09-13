// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionAssets struct {
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#config WorkerVersion#config}
	Config *WorkerVersionAssetsConfig `field:"optional" json:"config" yaml:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#jwt WorkerVersion#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
}

