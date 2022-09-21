// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccountConfig struct {
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
	// The name of the account that is displayed in the Cloudflare dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/account#name Account#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether 2FA is enforced on the account. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/account#enforce_twofactor Account#enforce_twofactor}
	EnforceTwofactor interface{} `field:"optional" json:"enforceTwofactor" yaml:"enforceTwofactor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/account#id Account#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Account type. Available values: `enterprise`, `standard`. Defaults to `standard`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/account#type Account#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

