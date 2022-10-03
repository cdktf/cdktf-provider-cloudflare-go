package pagesproject


type PagesProjectDeploymentConfigsPreview struct {
	// Compatibility date used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#compatibility_date PagesProject#compatibility_date}
	CompatibilityDate *string `field:"optional" json:"compatibilityDate" yaml:"compatibilityDate"`
	// Compatibility flags used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#compatibility_flags PagesProject#compatibility_flags}
	CompatibilityFlags *[]*string `field:"optional" json:"compatibilityFlags" yaml:"compatibilityFlags"`
	// D1 Databases used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#d1_databases PagesProject#d1_databases}
	D1Databases *map[string]*string `field:"optional" json:"d1Databases" yaml:"d1Databases"`
	// Durable Object namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#durable_object_namespaces PagesProject#durable_object_namespaces}
	DurableObjectNamespaces *map[string]*string `field:"optional" json:"durableObjectNamespaces" yaml:"durableObjectNamespaces"`
	// Environment variables for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#environment_variables PagesProject#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// KV namespaces used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#kv_namespaces PagesProject#kv_namespaces}
	KvNamespaces *map[string]*string `field:"optional" json:"kvNamespaces" yaml:"kvNamespaces"`
	// R2 Buckets used for Pages Functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/pages_project#r2_buckets PagesProject#r2_buckets}
	R2Buckets *map[string]*string `field:"optional" json:"r2Buckets" yaml:"r2Buckets"`
}

