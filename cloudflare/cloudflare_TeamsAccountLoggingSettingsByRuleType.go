// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type TeamsAccountLoggingSettingsByRuleType struct {
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#dns TeamsAccount#dns}
	Dns *TeamsAccountLoggingSettingsByRuleTypeDns `field:"required" json:"dns" yaml:"dns"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#http TeamsAccount#http}
	Http *TeamsAccountLoggingSettingsByRuleTypeHttp `field:"required" json:"http" yaml:"http"`
	// l4 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/teams_account#l4 TeamsAccount#l4}
	L4 *TeamsAccountLoggingSettingsByRuleTypeL4 `field:"required" json:"l4" yaml:"l4"`
}
