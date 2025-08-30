// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer#content_type LoadBalancer#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The http 'Location' header to include in the response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer#location LoadBalancer#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Text to include as the http body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer#message_body LoadBalancer#message_body}
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// The http status code to respond with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/resources/load_balancer#status_code LoadBalancer#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

