// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptAssets struct {
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#config WorkersScript#config}
	Config *WorkersScriptAssetsConfig `field:"optional" json:"config" yaml:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#jwt WorkersScript#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
}

