// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicySchedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Fridays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#fri ZeroTrustGatewayPolicy#fri}
	Fri *string `field:"optional" json:"fri" yaml:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Mondays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#mon ZeroTrustGatewayPolicy#mon}
	Mon *string `field:"optional" json:"mon" yaml:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Saturdays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#sat ZeroTrustGatewayPolicy#sat}
	Sat *string `field:"optional" json:"sat" yaml:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Sundays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#sun ZeroTrustGatewayPolicy#sun}
	Sun *string `field:"optional" json:"sun" yaml:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Thursdays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#thu ZeroTrustGatewayPolicy#thu}
	Thu *string `field:"optional" json:"thu" yaml:"thu"`
	// The time zone the rule will be evaluated against.
	//
	// If a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway will always use the current time at that time zone. If this parameter is omitted, then Gateway will use the time zone inferred from the user's source IP to evaluate the rule. If Gateway cannot determine the time zone from the IP, we will fall back to the time zone of the user's connected data center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#time_zone ZeroTrustGatewayPolicy#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Tuesdays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#tue ZeroTrustGatewayPolicy#tue}
	Tue *string `field:"optional" json:"tue" yaml:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing order from 00:00-24:00.
	//
	// If this parameter is omitted, the rule will be deactivated on Wednesdays.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_gateway_policy#wed ZeroTrustGatewayPolicy#wed}
	Wed *string `field:"optional" json:"wed" yaml:"wed"`
}

