// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptQueueBinding struct {
	// The name of the global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/workers_script#binding WorkersScript#binding}
	Binding *string `field:"required" json:"binding" yaml:"binding"`
	// Name of the queue you want to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.45.0/docs/resources/workers_script#queue WorkersScript#queue}
	Queue *string `field:"required" json:"queue" yaml:"queue"`
}

