// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RulesetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Type of Ruleset to create. Available values: `custom`, `managed`, `root`, `schema`, `zone`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#kind Ruleset#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the ruleset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#name Ruleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Point in the request/response lifecycle where the ruleset will be created.
	//
	// Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`, `http_config_settings`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#phase Ruleset#phase}
	Phase *string `field:"required" json:"phase" yaml:"phase"`
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#account_id Ruleset#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Brief summary of the ruleset and its intended use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#description Ruleset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#id Ruleset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rules Ruleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Name of entitlement that is shareable between entities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#shareable_entitlement_name Ruleset#shareable_entitlement_name}
	ShareableEntitlementName *string `field:"optional" json:"shareableEntitlementName" yaml:"shareableEntitlementName"`
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#zone_id Ruleset#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

