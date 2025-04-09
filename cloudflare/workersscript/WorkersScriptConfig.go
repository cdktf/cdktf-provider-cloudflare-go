// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersScriptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#account_id WorkersScript#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Module or Service Worker contents of the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#content WorkersScript#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Name of the script, used in URLs and route configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#script_name WorkersScript#script_name}
	ScriptName *string `field:"required" json:"scriptName" yaml:"scriptName"`
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#assets WorkersScript#assets}
	Assets *WorkersScriptAssets `field:"optional" json:"assets" yaml:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#bindings WorkersScript#bindings}
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
	// Name of the part in the multipart request that contains the script (e.g. the file adding a listener to the `fetch` event). Indicates a `service worker syntax` Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#body_part WorkersScript#body_part}
	BodyPart *string `field:"optional" json:"bodyPart" yaml:"bodyPart"`
	// Date indicating targeted support in the Workers runtime.
	//
	// Backwards incompatible fixes to the runtime following this date will not affect this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#compatibility_date WorkersScript#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Flags that enable or disable certain features in the Workers runtime.
	//
	// Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#compatibility_flags WorkersScript#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu of providing a completion token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#keep_assets WorkersScript#keep_assets}
	KeepAssets interface{} `field:"optional" json:"keepAssets" yaml:"keepAssets"`
	// List of binding types to keep from previous_upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#keep_bindings WorkersScript#keep_bindings}
	KeepBindings *[]*string `field:"optional" json:"keepBindings" yaml:"keepBindings"`
	// Whether Logpush is turned on for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#logpush WorkersScript#logpush}
	Logpush interface{} `field:"optional" json:"logpush" yaml:"logpush"`
	// Name of the part in the multipart request that contains the main module (e.g. the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#main_module WorkersScript#main_module}
	MainModule *string `field:"optional" json:"mainModule" yaml:"mainModule"`
	// Migrations to apply for Durable Objects associated with this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#migrations WorkersScript#migrations}
	Migrations *WorkersScriptMigrations `field:"optional" json:"migrations" yaml:"migrations"`
	// Observability settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#observability WorkersScript#observability}
	Observability *WorkersScriptObservability `field:"optional" json:"observability" yaml:"observability"`
	// Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#placement WorkersScript#placement}
	Placement *WorkersScriptPlacement `field:"optional" json:"placement" yaml:"placement"`
	// List of Workers that will consume logs from the attached Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#tail_consumers WorkersScript#tail_consumers}
	TailConsumers interface{} `field:"optional" json:"tailConsumers" yaml:"tailConsumers"`
	// Usage model for the Worker invocations. Available values: "standard".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/workers_script#usage_model WorkersScript#usage_model}
	UsageModel *string `field:"optional" json:"usageModel" yaml:"usageModel"`
}

