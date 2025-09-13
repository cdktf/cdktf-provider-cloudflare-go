// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerVersionConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#account_id WorkerVersion#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#worker_id WorkerVersion#worker_id}
	WorkerId *string `field:"required" json:"workerId" yaml:"workerId"`
	// Metadata about the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#annotations WorkerVersion#annotations}
	Annotations *WorkerVersionAnnotations `field:"optional" json:"annotations" yaml:"annotations"`
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#assets WorkerVersion#assets}
	Assets *WorkerVersionAssets `field:"optional" json:"assets" yaml:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#bindings WorkerVersion#bindings}
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
	// Date indicating targeted support in the Workers runtime.
	//
	// Backwards incompatible fixes to the runtime following this date will not affect this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#compatibility_date WorkerVersion#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Flags that enable or disable certain features in the Workers runtime.
	//
	// Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#compatibility_flags WorkerVersion#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// Resource limits enforced at runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#limits WorkerVersion#limits}
	Limits *WorkerVersionLimits `field:"optional" json:"limits" yaml:"limits"`
	// The name of the main module in the `modules` array (e.g. the name of the module that exports a `fetch` handler).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#main_module WorkerVersion#main_module}
	MainModule *string `field:"optional" json:"mainModule" yaml:"mainModule"`
	// Migrations for Durable Objects associated with the version. Migrations are applied when the version is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#migrations WorkerVersion#migrations}
	Migrations *WorkerVersionMigrations `field:"optional" json:"migrations" yaml:"migrations"`
	// Code, sourcemaps, and other content used at runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#modules WorkerVersion#modules}
	Modules interface{} `field:"optional" json:"modules" yaml:"modules"`
	// Placement settings for the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#placement WorkerVersion#placement}
	Placement *WorkerVersionPlacement `field:"optional" json:"placement" yaml:"placement"`
	// Usage model for the version. Available values: "standard", "bundled", "unbound".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/worker_version#usage_model WorkerVersion#usage_model}
	UsageModel *string `field:"optional" json:"usageModel" yaml:"usageModel"`
}

