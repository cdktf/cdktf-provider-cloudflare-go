// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdlpdataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDlpDatasetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#account_id ZeroTrustDlpDataset#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#name ZeroTrustDlpDataset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Only applies to custom word lists.
	//
	// Determines if the words should be matched in a case-sensitive manner
	// Cannot be set to false if `secret` is true or undefined
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#case_sensitive ZeroTrustDlpDataset#case_sensitive}
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#dataset_id ZeroTrustDlpDataset#dataset_id}.
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The description of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#description ZeroTrustDlpDataset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Dataset encoding version.
	//
	// Non-secret custom word lists with no header are always version 1.
	// Secret EDM lists with no header are version 1.
	// Multicolumn CSV with headers are version 2.
	// Omitting this field provides the default value 0, which is interpreted
	// the same as 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#encoding_version ZeroTrustDlpDataset#encoding_version}
	EncodingVersion *float64 `field:"optional" json:"encodingVersion" yaml:"encodingVersion"`
	// Generate a secret dataset.
	//
	// If true, the response will include a secret to use with the EDM encoder.
	// If false, the response has no secret and the dataset is uploaded in plaintext.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/zero_trust_dlp_dataset#secret ZeroTrustDlpDataset#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

