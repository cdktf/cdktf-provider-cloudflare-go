// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptBindings struct {
	// A JavaScript variable name for the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of resource that the binding provides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#type WorkersScript#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// R2 bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#bucket_name WorkersScript#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Identifier of the certificate to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#certificate_id WorkersScript#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// The exported class name of the Durable Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#class_name WorkersScript#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The dataset name to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#dataset WorkersScript#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The environment of the script_name to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#environment WorkersScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#id WorkersScript#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the Vectorize index to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#index_name WorkersScript#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// JSON data to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#json WorkersScript#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// Namespace to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#namespace WorkersScript#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Namespace identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#namespace_id WorkersScript#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#outbound WorkersScript#outbound}
	Outbound *WorkersScriptBindingsOutbound `field:"optional" json:"outbound" yaml:"outbound"`
	// Name of the Queue to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#queue_name WorkersScript#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The script where the Durable Object is defined, if it is external to this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#script_name WorkersScript#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Name of Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#service WorkersScript#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// The text value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.0.0/docs/resources/workers_script#text WorkersScript#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

