// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type AccessPolicyIncludeAzure struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#id AccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/access_policy#identity_provider_id AccessPolicy#identity_provider_id}.
	IdentityProviderId *string `field:"optional" json:"identityProviderId" yaml:"identityProviderId"`
}

