// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type HealthcheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/healthcheck#create Healthcheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

