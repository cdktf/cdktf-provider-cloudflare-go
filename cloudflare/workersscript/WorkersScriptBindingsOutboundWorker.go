// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptBindingsOutboundWorker struct {
	// Environment of the outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script#environment WorkersScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Name of the outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/workers_script#service WorkersScript#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

