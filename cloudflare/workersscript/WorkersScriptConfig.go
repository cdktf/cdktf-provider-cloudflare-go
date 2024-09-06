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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#account_id WorkersScript#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The script content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#content WorkersScript#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name for the script. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// analytics_engine_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#analytics_engine_binding WorkersScript#analytics_engine_binding}
	AnalyticsEngineBinding interface{} `field:"optional" json:"analyticsEngineBinding" yaml:"analyticsEngineBinding"`
	// The date to use for the compatibility flag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#compatibility_date WorkersScript#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Worker Scripts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#compatibility_flags WorkersScript#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// d1_database_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#d1_database_binding WorkersScript#d1_database_binding}
	D1DatabaseBinding interface{} `field:"optional" json:"d1DatabaseBinding" yaml:"d1DatabaseBinding"`
	// Name of the Workers for Platforms dispatch namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#dispatch_namespace WorkersScript#dispatch_namespace}
	DispatchNamespace *string `field:"optional" json:"dispatchNamespace" yaml:"dispatchNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#id WorkersScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kv_namespace_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#kv_namespace_binding WorkersScript#kv_namespace_binding}
	KvNamespaceBinding interface{} `field:"optional" json:"kvNamespaceBinding" yaml:"kvNamespaceBinding"`
	// Enabling allows Worker events to be sent to a defined Logpush destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#logpush WorkersScript#logpush}
	Logpush interface{} `field:"optional" json:"logpush" yaml:"logpush"`
	// Whether to upload Worker as a module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#module WorkersScript#module}
	Module interface{} `field:"optional" json:"module" yaml:"module"`
	// placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#placement WorkersScript#placement}
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// plain_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#plain_text_binding WorkersScript#plain_text_binding}
	PlainTextBinding interface{} `field:"optional" json:"plainTextBinding" yaml:"plainTextBinding"`
	// queue_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#queue_binding WorkersScript#queue_binding}
	QueueBinding interface{} `field:"optional" json:"queueBinding" yaml:"queueBinding"`
	// r2_bucket_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#r2_bucket_binding WorkersScript#r2_bucket_binding}
	R2BucketBinding interface{} `field:"optional" json:"r2BucketBinding" yaml:"r2BucketBinding"`
	// secret_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#secret_text_binding WorkersScript#secret_text_binding}
	SecretTextBinding interface{} `field:"optional" json:"secretTextBinding" yaml:"secretTextBinding"`
	// service_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#service_binding WorkersScript#service_binding}
	ServiceBinding interface{} `field:"optional" json:"serviceBinding" yaml:"serviceBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#tags WorkersScript#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// webassembly_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/workers_script#webassembly_binding WorkersScript#webassembly_binding}
	WebassemblyBinding interface{} `field:"optional" json:"webassemblyBinding" yaml:"webassemblyBinding"`
}

