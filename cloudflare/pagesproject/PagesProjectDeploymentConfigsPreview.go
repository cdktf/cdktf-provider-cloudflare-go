// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagesproject


type PagesProjectDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#ai_bindings PagesProject#ai_bindings}
	AiBindings interface{} `field:"optional" json:"aiBindings" yaml:"aiBindings"`
	// Analytics Engine bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#analytics_engine_datasets PagesProject#analytics_engine_datasets}
	AnalyticsEngineDatasets interface{} `field:"optional" json:"analyticsEngineDatasets" yaml:"analyticsEngineDatasets"`
	// Browser bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#browsers PagesProject#browsers}
	Browsers interface{} `field:"optional" json:"browsers" yaml:"browsers"`
	// Compatibility date used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#compatibility_date PagesProject#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#compatibility_flags PagesProject#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// D1 databases used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#d1_databases PagesProject#d1_databases}
	D1Databases interface{} `field:"optional" json:"d1Databases" yaml:"d1Databases"`
	// Durable Object namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#durable_object_namespaces PagesProject#durable_object_namespaces}
	DurableObjectNamespaces interface{} `field:"optional" json:"durableObjectNamespaces" yaml:"durableObjectNamespaces"`
	// Environment variables used for builds and Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#env_vars PagesProject#env_vars}
	EnvVars interface{} `field:"optional" json:"envVars" yaml:"envVars"`
	// Hyperdrive bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#hyperdrive_bindings PagesProject#hyperdrive_bindings}
	HyperdriveBindings interface{} `field:"optional" json:"hyperdriveBindings" yaml:"hyperdriveBindings"`
	// KV namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#kv_namespaces PagesProject#kv_namespaces}
	KvNamespaces interface{} `field:"optional" json:"kvNamespaces" yaml:"kvNamespaces"`
	// mTLS bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#mtls_certificates PagesProject#mtls_certificates}
	MtlsCertificates interface{} `field:"optional" json:"mtlsCertificates" yaml:"mtlsCertificates"`
	// Placement setting used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#placement PagesProject#placement}
	Placement *PagesProjectDeploymentConfigsPreviewPlacement `field:"optional" json:"placement" yaml:"placement"`
	// Queue Producer bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#queue_producers PagesProject#queue_producers}
	QueueProducers interface{} `field:"optional" json:"queueProducers" yaml:"queueProducers"`
	// R2 buckets used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#r2_buckets PagesProject#r2_buckets}
	R2Buckets interface{} `field:"optional" json:"r2Buckets" yaml:"r2Buckets"`
	// Services used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#services PagesProject#services}
	Services interface{} `field:"optional" json:"services" yaml:"services"`
	// Vectorize bindings used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/pages_project#vectorize_bindings PagesProject#vectorize_bindings}
	VectorizeBindings interface{} `field:"optional" json:"vectorizeBindings" yaml:"vectorizeBindings"`
}

