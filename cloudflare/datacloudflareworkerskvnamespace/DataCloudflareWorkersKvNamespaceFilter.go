// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflareworkerskvnamespace


type DataCloudflareWorkersKvNamespaceFilter struct {
	// Direction to order namespaces. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/workers_kv_namespace#direction DataCloudflareWorkersKvNamespace#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Field to order results by. Available values: "id", "title".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/data-sources/workers_kv_namespace#order DataCloudflareWorkersKvNamespace#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

