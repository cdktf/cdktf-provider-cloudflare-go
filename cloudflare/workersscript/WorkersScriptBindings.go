// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptBindings struct {
	// A JavaScript variable name for the binding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The kind of resource that the binding provides.
	//
	// Available values: "ai", "analytics_engine", "assets", "browser", "d1", "dispatch_namespace", "durable_object_namespace", "hyperdrive", "json", "kv_namespace", "mtls_certificate", "plain_text", "pipelines", "queue", "r2_bucket", "secret_text", "service", "tail_consumer", "vectorize", "version_metadata", "secrets_store_secret", "secret_key".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#type WorkersScript#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#algorithm WorkersScript#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// R2 bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#bucket_name WorkersScript#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Identifier of the certificate to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#certificate_id WorkersScript#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// The exported class name of the Durable Object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#class_name WorkersScript#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The name of the dataset to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#dataset WorkersScript#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The environment of the script_name to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#environment WorkersScript#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format). Available values: "raw", "pkcs8", "spki", "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#format WorkersScript#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Identifier of the D1 database to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#id WorkersScript#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the Vectorize index to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#index_name WorkersScript#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// JSON data to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#json WorkersScript#json}
	Json *string `field:"optional" json:"json" yaml:"json"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#key_base64 WorkersScript#key_base64}
	KeyBase64 *string `field:"optional" json:"keyBase64" yaml:"keyBase64"`
	// Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is "jwk".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#key_jwk WorkersScript#key_jwk}
	KeyJwk *string `field:"optional" json:"keyJwk" yaml:"keyJwk"`
	// Namespace to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#namespace WorkersScript#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Namespace identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#namespace_id WorkersScript#namespace_id}
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#outbound WorkersScript#outbound}
	Outbound *WorkersScriptBindingsOutbound `field:"optional" json:"outbound" yaml:"outbound"`
	// Name of the Pipeline to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#pipeline WorkersScript#pipeline}
	Pipeline *string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// Name of the Queue to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#queue_name WorkersScript#queue_name}
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The script where the Durable Object is defined, if it is external to this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#script_name WorkersScript#script_name}
	ScriptName *string `field:"optional" json:"scriptName" yaml:"scriptName"`
	// Name of the secret in the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#secret_name WorkersScript#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Name of Worker to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#service WorkersScript#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// ID of the store containing the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#store_id WorkersScript#store_id}
	StoreId *string `field:"optional" json:"storeId" yaml:"storeId"`
	// The text value to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#text WorkersScript#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#usages WorkersScript#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
	// Name of the Workflow to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#workflow_name WorkersScript#workflow_name}
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

