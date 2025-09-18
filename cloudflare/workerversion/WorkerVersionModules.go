// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionModules struct {
	// The file path of the module content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#content_file WorkerVersion#content_file}
	ContentFile *string `field:"required" json:"contentFile" yaml:"contentFile"`
	// The content type of the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#content_type WorkerVersion#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The name of the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

