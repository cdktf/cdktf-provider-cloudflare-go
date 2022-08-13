// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RateLimitCorrelate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/rate_limit#by RateLimit#by}.
	By *string `field:"optional" json:"by" yaml:"by"`
}

