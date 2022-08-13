// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type LoadBalancerPoolOriginSteering struct {
	// Defaults to `random`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer_pool#policy LoadBalancerPool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

