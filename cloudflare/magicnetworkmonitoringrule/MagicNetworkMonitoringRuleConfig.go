// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magicnetworkmonitoringrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicNetworkMonitoringRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#account_id MagicNetworkMonitoringRule#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The name of the rule.
	//
	// Must be unique. Supports characters A-Z, a-z, 0-9, underscore (_), dash (-), period (.), and tilde (~). You can’t have a space in the rule name. Max 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#name MagicNetworkMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP Prefixes within the rule via Magic Transit when the rule is triggered.
	//
	// Only available for users of Magic Transit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#automatic_advertisement MagicNetworkMonitoringRule#automatic_advertisement}
	AutomaticAdvertisement interface{} `field:"optional" json:"automaticAdvertisement" yaml:"automaticAdvertisement"`
	// The number of bits per second for the rule.
	//
	// When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#bandwidth MagicNetworkMonitoringRule#bandwidth}
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// The amount of time that the rule threshold must be exceeded to send an alert notification.
	//
	// The final value must be equivalent to one of the following 8 values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	// Available values: "1m", "5m", "10m", "15m", "20m", "30m", "45m", "60m".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#duration MagicNetworkMonitoringRule#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The number of packets per second for the rule.
	//
	// When this value is exceeded for the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#packet_threshold MagicNetworkMonitoringRule#packet_threshold}
	PacketThreshold *float64 `field:"optional" json:"packetThreshold" yaml:"packetThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_network_monitoring_rule#prefixes MagicNetworkMonitoringRule#prefixes}.
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

