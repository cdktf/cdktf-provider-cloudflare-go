package workerscript

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkerScriptConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The script content.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#content WorkerScript#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name for the script. **Modifying this attribute will force creation of a new resource.**.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#id WorkerScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kv_namespace_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#kv_namespace_binding WorkerScript#kv_namespace_binding}
	KvNamespaceBinding interface{} `field:"optional" json:"kvNamespaceBinding" yaml:"kvNamespaceBinding"`
	// Whether to upload Worker as a module.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#module WorkerScript#module}
	Module interface{} `field:"optional" json:"module" yaml:"module"`
	// plain_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#plain_text_binding WorkerScript#plain_text_binding}
	PlainTextBinding interface{} `field:"optional" json:"plainTextBinding" yaml:"plainTextBinding"`
	// r2_bucket_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#r2_bucket_binding WorkerScript#r2_bucket_binding}
	R2BucketBinding interface{} `field:"optional" json:"r2BucketBinding" yaml:"r2BucketBinding"`
	// secret_text_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#secret_text_binding WorkerScript#secret_text_binding}
	SecretTextBinding interface{} `field:"optional" json:"secretTextBinding" yaml:"secretTextBinding"`
	// service_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#service_binding WorkerScript#service_binding}
	ServiceBinding interface{} `field:"optional" json:"serviceBinding" yaml:"serviceBinding"`
	// webassembly_binding block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#webassembly_binding WorkerScript#webassembly_binding}
	WebassemblyBinding interface{} `field:"optional" json:"webassemblyBinding" yaml:"webassemblyBinding"`
}

