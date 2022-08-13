// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type WorkerScriptServiceBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#service WorkerScript#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#environment WorkerScript#environment}.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
}

