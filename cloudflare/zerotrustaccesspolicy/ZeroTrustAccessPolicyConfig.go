// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Defines the action Access will take if the policy matches the user. Available values: `allow`, `deny`, `non_identity`, `bypass`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#decision ZeroTrustAccessPolicy#decision}
	Decision *string `field:"required" json:"decision" yaml:"decision"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#include ZeroTrustAccessPolicy#include}
	Include interface{} `field:"required" json:"include" yaml:"include"`
	// Friendly name of the Access Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#name ZeroTrustAccessPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#account_id ZeroTrustAccessPolicy#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The ID of the application the policy is associated with.
	//
	// Required when using `precedence`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#application_id ZeroTrustAccessPolicy#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// approval_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#approval_group ZeroTrustAccessPolicy#approval_group}
	ApprovalGroup interface{} `field:"optional" json:"approvalGroup" yaml:"approvalGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#approval_required ZeroTrustAccessPolicy#approval_required}.
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
	// connection_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#connection_rules ZeroTrustAccessPolicy#connection_rules}
	ConnectionRules *ZeroTrustAccessPolicyConnectionRules `field:"optional" json:"connectionRules" yaml:"connectionRules"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#exclude ZeroTrustAccessPolicy#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#id ZeroTrustAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Require this application to be served in an isolated browser for users matching this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#isolation_required ZeroTrustAccessPolicy#isolation_required}
	IsolationRequired interface{} `field:"optional" json:"isolationRequired" yaml:"isolationRequired"`
	// The unique precedence for policies on a single application. Required when using `application_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#precedence ZeroTrustAccessPolicy#precedence}
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The prompt to display to the user for a justification for accessing the resource. Required when using `purpose_justification_required`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#purpose_justification_prompt ZeroTrustAccessPolicy#purpose_justification_prompt}
	PurposeJustificationPrompt *string `field:"optional" json:"purposeJustificationPrompt" yaml:"purposeJustificationPrompt"`
	// Whether to prompt the user for a justification for accessing the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#purpose_justification_required ZeroTrustAccessPolicy#purpose_justification_required}
	PurposeJustificationRequired interface{} `field:"optional" json:"purposeJustificationRequired" yaml:"purposeJustificationRequired"`
	// require block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#require ZeroTrustAccessPolicy#require}
	Require interface{} `field:"optional" json:"require" yaml:"require"`
	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#session_duration ZeroTrustAccessPolicy#session_duration}
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.52.0/docs/resources/zero_trust_access_policy#zone_id ZeroTrustAccessPolicy#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

