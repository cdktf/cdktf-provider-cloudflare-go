// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type ApiTokenPolicy struct {
	// List of permissions groups IDs. See [documentation](https://developers.cloudflare.com/api/tokens/create/permissions) for more information.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#permission_groups ApiToken#permission_groups}
	PermissionGroups *[]*string `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// Describes what operations against which resources are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#resources ApiToken#resources}
	Resources *map[string]*string `field:"required" json:"resources" yaml:"resources"`
	// Effect of the policy. Available values: `allow`, `deny`. Defaults to `allow`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#effect ApiToken#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
}
