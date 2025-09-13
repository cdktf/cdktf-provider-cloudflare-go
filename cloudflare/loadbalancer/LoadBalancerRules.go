// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadBalancerRules struct {
	// The condition expressions to evaluate.
	//
	// If the condition evaluates to true, the overrides or fixed_response in this rule will be applied. An empty condition is always true. For more details on condition expressions, please see https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#condition LoadBalancer#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#disabled LoadBalancer#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of routing to a pool.
	//
	// If a fixed_response is supplied the rule will be marked as terminates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#fixed_response LoadBalancer#fixed_response}
	FixedResponse *LoadBalancerRulesFixedResponse `field:"optional" json:"fixedResponse" yaml:"fixedResponse"`
	// Name of this rule. Only used for human readability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#name LoadBalancer#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A collection of overrides to apply to the load balancer when this rule's condition is true.
	//
	// All fields are optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#overrides LoadBalancer#overrides}
	Overrides *LoadBalancerRulesOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// The order in which rules should be executed in relation to each other.
	//
	// Lower values are executed first. Values do not need to be sequential. If no value is provided for any rule the array order of the rules field will be used to assign a priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#priority LoadBalancer#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after processing this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/load_balancer#terminates LoadBalancer#terminates}
	Terminates interface{} `field:"optional" json:"terminates" yaml:"terminates"`
}

