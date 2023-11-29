// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerscript


type WorkerScriptPlainTextBinding struct {
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The plain text you want to store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.20.0/docs/resources/worker_script#text WorkerScript#text}
	Text *string `field:"required" json:"text" yaml:"text"`
}

