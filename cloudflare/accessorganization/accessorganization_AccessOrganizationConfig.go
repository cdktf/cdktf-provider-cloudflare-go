package accessorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessOrganizationConfig struct {
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
	// The unique subdomain assigned to your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#auth_domain AccessOrganization#auth_domain}
	AuthDomain *string `field:"required" json:"authDomain" yaml:"authDomain"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#account_id AccessOrganization#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#id AccessOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#is_ui_read_only AccessOrganization#is_ui_read_only}
	IsUiReadOnly interface{} `field:"optional" json:"isUiReadOnly" yaml:"isUiReadOnly"`
	// login_design block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#login_design AccessOrganization#login_design}
	LoginDesign interface{} `field:"optional" json:"loginDesign" yaml:"loginDesign"`
	// The name of your Zero Trust organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#name AccessOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_organization#zone_id AccessOrganization#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

