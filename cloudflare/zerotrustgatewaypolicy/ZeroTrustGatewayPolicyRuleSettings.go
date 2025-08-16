// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs.
	//
	// Keys are header names, pointing to an array with its header value(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#add_headers ZeroTrustGatewayPolicy#add_headers}
	AddHeaders interface{} `field:"optional" json:"addHeaders" yaml:"addHeaders"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#allow_child_bypass ZeroTrustGatewayPolicy#allow_child_bypass}
	AllowChildBypass interface{} `field:"optional" json:"allowChildBypass" yaml:"allowChildBypass"`
	// Settings for the Audit SSH action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#audit_ssh ZeroTrustGatewayPolicy#audit_ssh}
	AuditSsh *ZeroTrustGatewayPolicyRuleSettingsAuditSsh `field:"optional" json:"auditSsh" yaml:"auditSsh"`
	// Configure how browser isolation behaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#biso_admin_controls ZeroTrustGatewayPolicy#biso_admin_controls}
	BisoAdminControls *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls `field:"optional" json:"bisoAdminControls" yaml:"bisoAdminControls"`
	// Custom block page settings. If missing/null, blocking will use the the account settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#block_page ZeroTrustGatewayPolicy#block_page}
	BlockPage *ZeroTrustGatewayPolicyRuleSettingsBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// Enable the custom block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#block_page_enabled ZeroTrustGatewayPolicy#block_page_enabled}
	BlockPageEnabled interface{} `field:"optional" json:"blockPageEnabled" yaml:"blockPageEnabled"`
	// The text describing why this block occurred, displayed on the custom block page (if enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#block_reason ZeroTrustGatewayPolicy#block_reason}
	BlockReason *string `field:"optional" json:"blockReason" yaml:"blockReason"`
	// Set by children MSP accounts to bypass their parent's rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#bypass_parent_rule ZeroTrustGatewayPolicy#bypass_parent_rule}
	BypassParentRule interface{} `field:"optional" json:"bypassParentRule" yaml:"bypassParentRule"`
	// Configure how session check behaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#check_session ZeroTrustGatewayPolicy#check_session}
	CheckSession *ZeroTrustGatewayPolicyRuleSettingsCheckSession `field:"optional" json:"checkSession" yaml:"checkSession"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	//
	// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally' are set. DNS queries will route to the address closest to their origin. Only valid when a rule's action is set to 'resolve'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#dns_resolvers ZeroTrustGatewayPolicy#dns_resolvers}
	DnsResolvers *ZeroTrustGatewayPolicyRuleSettingsDnsResolvers `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// Configure how Gateway Proxy traffic egresses.
	//
	// You can enable this setting for rules with Egress actions and filters, or omit it to indicate local egress via WARP IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#egress ZeroTrustGatewayPolicy#egress}
	Egress *ZeroTrustGatewayPolicyRuleSettingsEgress `field:"optional" json:"egress" yaml:"egress"`
	// Set to true, to ignore the category matches at CNAME domains in a response.
	//
	// If unchecked, the categories in this rule will be checked against all the CNAME domain categories in a response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#ignore_cname_category_matches ZeroTrustGatewayPolicy#ignore_cname_category_matches}
	IgnoreCnameCategoryMatches interface{} `field:"optional" json:"ignoreCnameCategoryMatches" yaml:"ignoreCnameCategoryMatches"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#insecure_disable_dnssec_validation ZeroTrustGatewayPolicy#insecure_disable_dnssec_validation}
	InsecureDisableDnssecValidation interface{} `field:"optional" json:"insecureDisableDnssecValidation" yaml:"insecureDisableDnssecValidation"`
	// Set to true to enable IPs in DNS resolver category blocks.
	//
	// By default categories only block based on domain names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#ip_categories ZeroTrustGatewayPolicy#ip_categories}
	IpCategories interface{} `field:"optional" json:"ipCategories" yaml:"ipCategories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks.
	//
	// By default indicator feeds only block based on domain names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#ip_indicator_feeds ZeroTrustGatewayPolicy#ip_indicator_feeds}
	IpIndicatorFeeds interface{} `field:"optional" json:"ipIndicatorFeeds" yaml:"ipIndicatorFeeds"`
	// Send matching traffic to the supplied destination IP address and port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#l4override ZeroTrustGatewayPolicy#l4override}
	L4Override *ZeroTrustGatewayPolicyRuleSettingsL4Override `field:"optional" json:"l4Override" yaml:"l4Override"`
	// Configure a notification to display on the user's device when this rule is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#notification_settings ZeroTrustGatewayPolicy#notification_settings}
	NotificationSettings *ZeroTrustGatewayPolicyRuleSettingsNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// Override matching DNS queries with a hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#override_host ZeroTrustGatewayPolicy#override_host}
	OverrideHost *string `field:"optional" json:"overrideHost" yaml:"overrideHost"`
	// Override matching DNS queries with an IP or set of IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#override_ips ZeroTrustGatewayPolicy#override_ips}
	OverrideIps *[]*string `field:"optional" json:"overrideIps" yaml:"overrideIps"`
	// Configure DLP payload logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#payload_log ZeroTrustGatewayPolicy#payload_log}
	PayloadLog *ZeroTrustGatewayPolicyRuleSettingsPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// Settings that apply to quarantine rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#quarantine ZeroTrustGatewayPolicy#quarantine}
	Quarantine *ZeroTrustGatewayPolicyRuleSettingsQuarantine `field:"optional" json:"quarantine" yaml:"quarantine"`
	// Settings that apply to redirect rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#redirect ZeroTrustGatewayPolicy#redirect}
	Redirect *ZeroTrustGatewayPolicyRuleSettingsRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// Configure to forward the query to the internal DNS service, passing the specified 'view_id' as input.
	//
	// Cannot be set when 'dns_resolvers' are specified or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is set to 'resolve'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#resolve_dns_internally ZeroTrustGatewayPolicy#resolve_dns_internally}
	ResolveDnsInternally *ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally `field:"optional" json:"resolveDnsInternally" yaml:"resolveDnsInternally"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot be set when 'dns_resolvers' are specified or 'resolve_dns_internally' is set. Only valid when a rule's action is set to 'resolve'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#resolve_dns_through_cloudflare ZeroTrustGatewayPolicy#resolve_dns_through_cloudflare}
	ResolveDnsThroughCloudflare interface{} `field:"optional" json:"resolveDnsThroughCloudflare" yaml:"resolveDnsThroughCloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.4/docs/resources/zero_trust_gateway_policy#untrusted_cert ZeroTrustGatewayPolicy#untrusted_cert}
	UntrustedCert *ZeroTrustGatewayPolicyRuleSettingsUntrustedCert `field:"optional" json:"untrustedCert" yaml:"untrustedCert"`
}

