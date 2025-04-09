// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccessapplication


type ZeroTrustAccessApplicationPolicies struct {
	// The rules that define how users may connect to the targets secured by your application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#connection_rules ZeroTrustAccessApplication#connection_rules}
	ConnectionRules *ZeroTrustAccessApplicationPoliciesConnectionRules `field:"optional" json:"connectionRules" yaml:"connectionRules"`
	// The action Access will take if a user matches this policy.
	//
	// Infrastructure application policies can only use the Allow action.
	// Available values: "allow", "deny", "non_identity", "bypass".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#decision ZeroTrustAccessApplication#decision}
	Decision *string `field:"optional" json:"decision" yaml:"decision"`
	// Rules evaluated with a NOT logical operator.
	//
	// To match the policy, a user cannot meet any of the Exclude rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#exclude ZeroTrustAccessApplication#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// The UUID of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#id ZeroTrustAccessApplication#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#include ZeroTrustAccessApplication#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// The name of the Access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#name ZeroTrustAccessApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The order of execution for this policy. Must be unique for each policy within an app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#precedence ZeroTrustAccessApplication#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// Rules evaluated with an AND logical operator.
	//
	// To match the policy, a user must meet all of the Require rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/zero_trust_access_application#require ZeroTrustAccessApplication#require}
	Require interface{} `field:"optional" json:"require" yaml:"require"`
}

