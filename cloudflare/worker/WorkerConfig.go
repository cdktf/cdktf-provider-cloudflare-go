// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#account_id Worker#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#name Worker#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether logpush is enabled for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#logpush Worker#logpush}
	Logpush interface{} `field:"optional" json:"logpush" yaml:"logpush"`
	// Observability settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#observability Worker#observability}
	Observability *WorkerObservability `field:"optional" json:"observability" yaml:"observability"`
	// Subdomain settings for the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#subdomain Worker#subdomain}
	Subdomain *WorkerSubdomain `field:"optional" json:"subdomain" yaml:"subdomain"`
	// Tags associated with the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#tags Worker#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Other Workers that should consume logs from the Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker#tail_consumers Worker#tail_consumers}
	TailConsumers interface{} `field:"optional" json:"tailConsumers" yaml:"tailConsumers"`
}

