// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type WorkerScriptR2BucketBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#bucket_name WorkerScript#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/worker_script#name WorkerScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

