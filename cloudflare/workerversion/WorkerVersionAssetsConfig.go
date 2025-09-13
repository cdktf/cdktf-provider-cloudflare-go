// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content. Available values: "auto-trailing-slash", "force-trailing-slash", "drop-trailing-slash", "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#html_handling WorkerVersion#html_handling}
	HtmlHandling *string `field:"optional" json:"htmlHandling" yaml:"htmlHandling"`
	// Determines the response when a request does not match a static asset, and there is no Worker script.
	//
	// Available values: "none", "404-page", "single-page-application".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#not_found_handling WorkerVersion#not_found_handling}
	NotFoundHandling *string `field:"optional" json:"notFoundHandling" yaml:"notFoundHandling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	//
	// Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#run_worker_first WorkerVersion#run_worker_first}
	RunWorkerFirst *[]*string `field:"optional" json:"runWorkerFirst" yaml:"runWorkerFirst"`
}

