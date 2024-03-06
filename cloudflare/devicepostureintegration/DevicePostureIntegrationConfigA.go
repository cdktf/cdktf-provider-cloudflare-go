// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devicepostureintegration


type DevicePostureIntegrationConfigA struct {
	// The Access client ID to be used as the `Cf-Access-Client-ID` header when making a request to the `api_url`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#access_client_id DevicePostureIntegration#access_client_id}
	AccessClientId *string `field:"optional" json:"accessClientId" yaml:"accessClientId"`
	// The Access client secret to be used as the `Cf-Access-Client-Secret` header when making a request to the `api_url`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#access_client_secret DevicePostureIntegration#access_client_secret}
	AccessClientSecret *string `field:"optional" json:"accessClientSecret" yaml:"accessClientSecret"`
	// The third-party API's URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#api_url DevicePostureIntegration#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// The third-party authorization API URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#auth_url DevicePostureIntegration#auth_url}
	AuthUrl *string `field:"optional" json:"authUrl" yaml:"authUrl"`
	// The client identifier for authenticating API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#client_id DevicePostureIntegration#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client key for authenticating API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#client_key DevicePostureIntegration#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// The client secret for authenticating API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#client_secret DevicePostureIntegration#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The customer identifier for authenticating API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.26.0/docs/resources/device_posture_integration#customer_id DevicePostureIntegration#customer_id}
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
}

