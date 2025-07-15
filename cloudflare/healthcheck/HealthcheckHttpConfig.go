// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcheck


type HealthcheckHttpConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#allow_insecure Healthcheck#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// A case-insensitive sub-string to look for in the response body.
	//
	// If this string is not found, the origin will be marked as unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#expected_body Healthcheck#expected_body}
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all codes starting with 2) of the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#expected_codes Healthcheck#expected_codes}
	ExpectedCodes *[]*string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Follow redirects if the origin returns a 3xx status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#follow_redirects Healthcheck#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// The HTTP request headers to send in the health check.
	//
	// It is recommended you set a Host header by default. The User-Agent header cannot be overridden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#header Healthcheck#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The HTTP method to use for the health check. Available values: "GET", "HEAD".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#method Healthcheck#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The endpoint path to health check against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#path Healthcheck#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port number to connect to for the health check.
	//
	// Defaults to 80 if type is HTTP or 443 if type is HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/healthcheck#port Healthcheck#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

