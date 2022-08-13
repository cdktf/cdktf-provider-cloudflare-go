// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevicePostureIntegrationConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#account_id DevicePostureIntegration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#name DevicePostureIntegration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#type DevicePostureIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#config DevicePostureIntegration#config}
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#id DevicePostureIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#identifier DevicePostureIntegration#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/device_posture_integration#interval DevicePostureIntegration#interval}.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
}

