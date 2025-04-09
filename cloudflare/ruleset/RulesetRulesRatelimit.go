// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be incremented.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#characteristics Ruleset#characteristics}
	Characteristics *[]*string `field:"required" json:"characteristics" yaml:"characteristics"`
	// Period in seconds over which the counter is being incremented. Available values: 10, 60, 600, 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#period Ruleset#period}
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Defines when the ratelimit counter should be incremented.
	//
	// It is optional and defaults to the same as the rule's expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#counting_expression Ruleset#counting_expression}
	CountingExpression *string `field:"optional" json:"countingExpression" yaml:"countingExpression"`
	// Period of time in seconds after which the action will be disabled following its first execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#mitigation_timeout Ruleset#mitigation_timeout}
	MitigationTimeout *float64 `field:"optional" json:"mitigationTimeout" yaml:"mitigationTimeout"`
	// The threshold of requests per period after which the action will be executed for the first time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#requests_per_period Ruleset#requests_per_period}
	RequestsPerPeriod *float64 `field:"optional" json:"requestsPerPeriod" yaml:"requestsPerPeriod"`
	// Defines if ratelimit counting is only done when an origin is reached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#requests_to_origin Ruleset#requests_to_origin}
	RequestsToOrigin interface{} `field:"optional" json:"requestsToOrigin" yaml:"requestsToOrigin"`
	// The score threshold per period for which the action will be executed the first time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#score_per_period Ruleset#score_per_period}
	ScorePerPeriod *float64 `field:"optional" json:"scorePerPeriod" yaml:"scorePerPeriod"`
	// The response header name provided by the origin which should contain the score to increment ratelimit counter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/ruleset#score_response_header_name Ruleset#score_response_header_name}
	ScoreResponseHeaderName *string `field:"optional" json:"scoreResponseHeaderName" yaml:"scoreResponseHeaderName"`
}

