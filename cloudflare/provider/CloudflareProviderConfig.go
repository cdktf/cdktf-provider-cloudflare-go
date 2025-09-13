// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type CloudflareProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#alias CloudflareProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The API key for operations.
	//
	// Alternatively, can be configured using the `CLOUDFLARE_API_KEY` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/fundamentals/api/get-started/keys/#limitations), API tokens should be used instead. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#api_key CloudflareProvider#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The API Token for operations.
	//
	// Alternatively, can be configured using the `CLOUDFLARE_API_TOKEN` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#api_token CloudflareProvider#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// A special Cloudflare API key good for a restricted set of endpoints.
	//
	// Alternatively, can be configured using the `CLOUDFLARE_API_USER_SERVICE_KEY` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#api_user_service_key CloudflareProvider#api_user_service_key}
	ApiUserServiceKey *string `field:"optional" json:"apiUserServiceKey" yaml:"apiUserServiceKey"`
	// Value to override the default HTTP client base URL. Alternatively, can be configured using the `base_url` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#base_url CloudflareProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// A registered Cloudflare email address.
	//
	// Alternatively, can be configured using the `CLOUDFLARE_EMAIL` environment variable. Required when using `api_key`. Conflicts with `api_token`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#email CloudflareProvider#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// A value to append to the HTTP User Agent for all API calls.
	//
	// This value is not something most users need to modify however, if you are using a non-standard provider or operator configuration, this is recommended to assist in uniquely identifying your traffic. **Setting this value will remove the Terraform version from the HTTP User Agent string and may have unintended consequences**. Alternatively, can be configured using the `CLOUDFLARE_USER_AGENT_OPERATOR_SUFFIX` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs#user_agent_operator_suffix CloudflareProvider#user_agent_operator_suffix}
	UserAgentOperatorSuffix *string `field:"optional" json:"userAgentOperatorSuffix" yaml:"userAgentOperatorSuffix"`
}

