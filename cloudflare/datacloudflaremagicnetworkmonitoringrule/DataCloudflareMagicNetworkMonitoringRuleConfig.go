// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacloudflaremagicnetworkmonitoringrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareMagicNetworkMonitoringRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/magic_network_monitoring_rule#account_id DataCloudflareMagicNetworkMonitoringRule#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The id of the rule. Must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.9.0/docs/data-sources/magic_network_monitoring_rule#rule_id DataCloudflareMagicNetworkMonitoringRule#rule_id}
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
}

