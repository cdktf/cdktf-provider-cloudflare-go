// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RulesetRulesActionParametersResponse struct {
	// Body content to include in the response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#content Ruleset#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// HTTP content type to send in the response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// HTTP status code to send in the response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}
