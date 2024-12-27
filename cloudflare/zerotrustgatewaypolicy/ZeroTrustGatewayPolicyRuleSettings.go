// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettings struct {
	// Add custom headers to allowed requests in the form of key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#add_headers ZeroTrustGatewayPolicy#add_headers}
	AddHeaders *map[string]*string `field:"optional" json:"addHeaders" yaml:"addHeaders"`
	// Allow parent MSP accounts to enable bypass their children's rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#allow_child_bypass ZeroTrustGatewayPolicy#allow_child_bypass}
	AllowChildBypass interface{} `field:"optional" json:"allowChildBypass" yaml:"allowChildBypass"`
	// audit_ssh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#audit_ssh ZeroTrustGatewayPolicy#audit_ssh}
	AuditSsh *ZeroTrustGatewayPolicyRuleSettingsAuditSsh `field:"optional" json:"auditSsh" yaml:"auditSsh"`
	// biso_admin_controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#biso_admin_controls ZeroTrustGatewayPolicy#biso_admin_controls}
	BisoAdminControls *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls `field:"optional" json:"bisoAdminControls" yaml:"bisoAdminControls"`
	// Indicator of block page enablement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#block_page_enabled ZeroTrustGatewayPolicy#block_page_enabled}
	BlockPageEnabled interface{} `field:"optional" json:"blockPageEnabled" yaml:"blockPageEnabled"`
	// The displayed reason for a user being blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#block_page_reason ZeroTrustGatewayPolicy#block_page_reason}
	BlockPageReason *string `field:"optional" json:"blockPageReason" yaml:"blockPageReason"`
	// Allow child MSP accounts to bypass their parent's rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#bypass_parent_rule ZeroTrustGatewayPolicy#bypass_parent_rule}
	BypassParentRule interface{} `field:"optional" json:"bypassParentRule" yaml:"bypassParentRule"`
	// check_session block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#check_session ZeroTrustGatewayPolicy#check_session}
	CheckSession *ZeroTrustGatewayPolicyRuleSettingsCheckSession `field:"optional" json:"checkSession" yaml:"checkSession"`
	// dns_resolvers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#dns_resolvers ZeroTrustGatewayPolicy#dns_resolvers}
	DnsResolvers *ZeroTrustGatewayPolicyRuleSettingsDnsResolvers `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#egress ZeroTrustGatewayPolicy#egress}
	Egress *ZeroTrustGatewayPolicyRuleSettingsEgress `field:"optional" json:"egress" yaml:"egress"`
	// Set to true, to ignore the category matches at CNAME domains in a response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#ignore_cname_category_matches ZeroTrustGatewayPolicy#ignore_cname_category_matches}
	IgnoreCnameCategoryMatches interface{} `field:"optional" json:"ignoreCnameCategoryMatches" yaml:"ignoreCnameCategoryMatches"`
	// Disable DNSSEC validation (must be Allow rule).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#insecure_disable_dnssec_validation ZeroTrustGatewayPolicy#insecure_disable_dnssec_validation}
	InsecureDisableDnssecValidation interface{} `field:"optional" json:"insecureDisableDnssecValidation" yaml:"insecureDisableDnssecValidation"`
	// Turns on IP category based filter on dns if the rule contains dns category checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#ip_categories ZeroTrustGatewayPolicy#ip_categories}
	IpCategories interface{} `field:"optional" json:"ipCategories" yaml:"ipCategories"`
	// l4override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#l4override ZeroTrustGatewayPolicy#l4override}
	L4Override *ZeroTrustGatewayPolicyRuleSettingsL4Override `field:"optional" json:"l4Override" yaml:"l4Override"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#notification_settings ZeroTrustGatewayPolicy#notification_settings}
	NotificationSettings *ZeroTrustGatewayPolicyRuleSettingsNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// The host to override matching DNS queries with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#override_host ZeroTrustGatewayPolicy#override_host}
	OverrideHost *string `field:"optional" json:"overrideHost" yaml:"overrideHost"`
	// The IPs to override matching DNS queries with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#override_ips ZeroTrustGatewayPolicy#override_ips}
	OverrideIps *[]*string `field:"optional" json:"overrideIps" yaml:"overrideIps"`
	// payload_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#payload_log ZeroTrustGatewayPolicy#payload_log}
	PayloadLog *ZeroTrustGatewayPolicyRuleSettingsPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// Enable sending queries that match the resolver policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot be set when `dns_resolvers` are specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#resolve_dns_through_cloudflare ZeroTrustGatewayPolicy#resolve_dns_through_cloudflare}
	ResolveDnsThroughCloudflare interface{} `field:"optional" json:"resolveDnsThroughCloudflare" yaml:"resolveDnsThroughCloudflare"`
	// untrusted_cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/zero_trust_gateway_policy#untrusted_cert ZeroTrustGatewayPolicy#untrusted_cert}
	UntrustedCert *ZeroTrustGatewayPolicyRuleSettingsUntrustedCert `field:"optional" json:"untrustedCert" yaml:"untrustedCert"`
}

