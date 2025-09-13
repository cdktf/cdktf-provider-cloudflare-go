// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionMigrationsStepsTransferredClasses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#from WorkerVersion#from}.
	From *string `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#from_script WorkerVersion#from_script}.
	FromScript *string `field:"optional" json:"fromScript" yaml:"fromScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#to WorkerVersion#to}.
	To *string `field:"optional" json:"to" yaml:"to"`
}

