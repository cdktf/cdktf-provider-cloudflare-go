// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptTailConsumers struct {
	// Name of Worker that is to be the consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#service WorkersScript#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional environment if the Worker utilizes one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#environment WorkersScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Optional dispatch namespace the script belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#namespace WorkersScript#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

