// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type DevicePostureRuleMatch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_rule#platform DevicePostureRule#platform}.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

