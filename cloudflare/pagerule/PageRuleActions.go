// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pagerule


type PageRuleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#always_use_https PageRule#always_use_https}.
	AlwaysUseHttps interface{} `field:"optional" json:"alwaysUseHttps" yaml:"alwaysUseHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#automatic_https_rewrites PageRule#automatic_https_rewrites}.
	AutomaticHttpsRewrites *string `field:"optional" json:"automaticHttpsRewrites" yaml:"automaticHttpsRewrites"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#browser_cache_ttl PageRule#browser_cache_ttl}.
	BrowserCacheTtl *float64 `field:"optional" json:"browserCacheTtl" yaml:"browserCacheTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#browser_check PageRule#browser_check}.
	BrowserCheck *string `field:"optional" json:"browserCheck" yaml:"browserCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#bypass_cache_on_cookie PageRule#bypass_cache_on_cookie}.
	BypassCacheOnCookie *string `field:"optional" json:"bypassCacheOnCookie" yaml:"bypassCacheOnCookie"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_by_device_type PageRule#cache_by_device_type}.
	CacheByDeviceType *string `field:"optional" json:"cacheByDeviceType" yaml:"cacheByDeviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_deception_armor PageRule#cache_deception_armor}.
	CacheDeceptionArmor *string `field:"optional" json:"cacheDeceptionArmor" yaml:"cacheDeceptionArmor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_key_fields PageRule#cache_key_fields}.
	CacheKeyFields *PageRuleActionsCacheKeyFields `field:"optional" json:"cacheKeyFields" yaml:"cacheKeyFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_level PageRule#cache_level}.
	CacheLevel *string `field:"optional" json:"cacheLevel" yaml:"cacheLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_on_cookie PageRule#cache_on_cookie}.
	CacheOnCookie *string `field:"optional" json:"cacheOnCookie" yaml:"cacheOnCookie"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#cache_ttl_by_status PageRule#cache_ttl_by_status}.
	CacheTtlByStatus *map[string]interface{} `field:"optional" json:"cacheTtlByStatus" yaml:"cacheTtlByStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#disable_apps PageRule#disable_apps}.
	DisableApps interface{} `field:"optional" json:"disableApps" yaml:"disableApps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#disable_performance PageRule#disable_performance}.
	DisablePerformance interface{} `field:"optional" json:"disablePerformance" yaml:"disablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#disable_security PageRule#disable_security}.
	DisableSecurity interface{} `field:"optional" json:"disableSecurity" yaml:"disableSecurity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#disable_zaraz PageRule#disable_zaraz}.
	DisableZaraz interface{} `field:"optional" json:"disableZaraz" yaml:"disableZaraz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#edge_cache_ttl PageRule#edge_cache_ttl}.
	EdgeCacheTtl *float64 `field:"optional" json:"edgeCacheTtl" yaml:"edgeCacheTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#email_obfuscation PageRule#email_obfuscation}.
	EmailObfuscation *string `field:"optional" json:"emailObfuscation" yaml:"emailObfuscation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#explicit_cache_control PageRule#explicit_cache_control}.
	ExplicitCacheControl *string `field:"optional" json:"explicitCacheControl" yaml:"explicitCacheControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#forwarding_url PageRule#forwarding_url}.
	ForwardingUrl *PageRuleActionsForwardingUrl `field:"optional" json:"forwardingUrl" yaml:"forwardingUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#host_header_override PageRule#host_header_override}.
	HostHeaderOverride *string `field:"optional" json:"hostHeaderOverride" yaml:"hostHeaderOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#ip_geolocation PageRule#ip_geolocation}.
	IpGeolocation *string `field:"optional" json:"ipGeolocation" yaml:"ipGeolocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#mirage PageRule#mirage}.
	Mirage *string `field:"optional" json:"mirage" yaml:"mirage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#opportunistic_encryption PageRule#opportunistic_encryption}.
	OpportunisticEncryption *string `field:"optional" json:"opportunisticEncryption" yaml:"opportunisticEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#origin_error_page_pass_thru PageRule#origin_error_page_pass_thru}.
	OriginErrorPagePassThru *string `field:"optional" json:"originErrorPagePassThru" yaml:"originErrorPagePassThru"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#polish PageRule#polish}.
	Polish *string `field:"optional" json:"polish" yaml:"polish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#resolve_override PageRule#resolve_override}.
	ResolveOverride *string `field:"optional" json:"resolveOverride" yaml:"resolveOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#respect_strong_etag PageRule#respect_strong_etag}.
	RespectStrongEtag *string `field:"optional" json:"respectStrongEtag" yaml:"respectStrongEtag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#response_buffering PageRule#response_buffering}.
	ResponseBuffering *string `field:"optional" json:"responseBuffering" yaml:"responseBuffering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#rocket_loader PageRule#rocket_loader}.
	RocketLoader *string `field:"optional" json:"rocketLoader" yaml:"rocketLoader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#security_level PageRule#security_level}.
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#sort_query_string_for_cache PageRule#sort_query_string_for_cache}.
	SortQueryStringForCache *string `field:"optional" json:"sortQueryStringForCache" yaml:"sortQueryStringForCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#ssl PageRule#ssl}.
	Ssl *string `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#true_client_ip_header PageRule#true_client_ip_header}.
	TrueClientIpHeader *string `field:"optional" json:"trueClientIpHeader" yaml:"trueClientIpHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/page_rule#waf PageRule#waf}.
	Waf *string `field:"optional" json:"waf" yaml:"waf"`
}

