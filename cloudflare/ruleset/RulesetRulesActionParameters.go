// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ruleset


type RulesetRulesActionParameters struct {
	// List of additional ports that caching can be enabled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#additional_cacheable_ports Ruleset#additional_cacheable_ports}
	AdditionalCacheablePorts *[]*float64 `field:"optional" json:"additionalCacheablePorts" yaml:"additionalCacheablePorts"`
	// Custom order for compression algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#algorithms Ruleset#algorithms}
	Algorithms interface{} `field:"optional" json:"algorithms" yaml:"algorithms"`
	// Turn on or off Automatic HTTPS Rewrites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#automatic_https_rewrites Ruleset#automatic_https_rewrites}
	AutomaticHttpsRewrites interface{} `field:"optional" json:"automaticHttpsRewrites" yaml:"automaticHttpsRewrites"`
	// Select which file extensions to minify automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#autominify Ruleset#autominify}
	Autominify *RulesetRulesActionParametersAutominify `field:"optional" json:"autominify" yaml:"autominify"`
	// Turn on or off Browser Integrity Check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#bic Ruleset#bic}
	Bic interface{} `field:"optional" json:"bic" yaml:"bic"`
	// Specify how long client browsers should cache the response.
	//
	// Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#browser_ttl Ruleset#browser_ttl}
	BrowserTtl *RulesetRulesActionParametersBrowserTtl `field:"optional" json:"browserTtl" yaml:"browserTtl"`
	// Mark whether the request’s response from origin is eligible for caching.
	//
	// Caching itself will still depend on the cache-control header and your other caching configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#cache Ruleset#cache}
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// Define which components of the request are included or excluded from the cache key Cloudflare uses to store the response in cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#cache_key Ruleset#cache_key}
	CacheKey *RulesetRulesActionParametersCacheKey `field:"optional" json:"cacheKey" yaml:"cacheKey"`
	// Mark whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#cache_reserve Ruleset#cache_reserve}
	CacheReserve *RulesetRulesActionParametersCacheReserve `field:"optional" json:"cacheReserve" yaml:"cacheReserve"`
	// Error response content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#content Ruleset#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content-type header to set with the response. Available values: "application/json", "text/xml", "text/plain", "text/html".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The cookie fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#cookie_fields Ruleset#cookie_fields}
	CookieFields interface{} `field:"optional" json:"cookieFields" yaml:"cookieFields"`
	// Turn off all active Cloudflare Apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#disable_apps Ruleset#disable_apps}
	DisableApps interface{} `field:"optional" json:"disableApps" yaml:"disableApps"`
	// Turn off Real User Monitoring (RUM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#disable_rum Ruleset#disable_rum}
	DisableRum interface{} `field:"optional" json:"disableRum" yaml:"disableRum"`
	// Turn off Zaraz.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#disable_zaraz Ruleset#disable_zaraz}
	DisableZaraz interface{} `field:"optional" json:"disableZaraz" yaml:"disableZaraz"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the Cloudflare edge network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#edge_ttl Ruleset#edge_ttl}
	EdgeTtl *RulesetRulesActionParametersEdgeTtl `field:"optional" json:"edgeTtl" yaml:"edgeTtl"`
	// Turn on or off Email Obfuscation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#email_obfuscation Ruleset#email_obfuscation}
	EmailObfuscation interface{} `field:"optional" json:"emailObfuscation" yaml:"emailObfuscation"`
	// Turn on or off Cloudflare Fonts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#fonts Ruleset#fonts}
	Fonts interface{} `field:"optional" json:"fonts" yaml:"fonts"`
	// Serve a redirect based on a bulk list lookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#from_list Ruleset#from_list}
	FromList *RulesetRulesActionParametersFromListStruct `field:"optional" json:"fromList" yaml:"fromList"`
	// Serve a redirect based on the request properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#from_value Ruleset#from_value}
	FromValue *RulesetRulesActionParametersFromValue `field:"optional" json:"fromValue" yaml:"fromValue"`
	// Map of request headers to modify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Rewrite the HTTP Host header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#host_header Ruleset#host_header}
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Turn on or off the Hotlink Protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#hotlink_protection Ruleset#hotlink_protection}
	HotlinkProtection interface{} `field:"optional" json:"hotlinkProtection" yaml:"hotlinkProtection"`
	// The ID of the ruleset to execute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Increment contains the delta to change the score and can be either positive or negative.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#increment Ruleset#increment}
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
	// The configuration to use for matched data logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#matched_data Ruleset#matched_data}
	MatchedData *RulesetRulesActionParametersMatchedData `field:"optional" json:"matchedData" yaml:"matchedData"`
	// Turn on or off Mirage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#mirage Ruleset#mirage}
	Mirage interface{} `field:"optional" json:"mirage" yaml:"mirage"`
	// Turn on or off Opportunistic Encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#opportunistic_encryption Ruleset#opportunistic_encryption}
	OpportunisticEncryption interface{} `field:"optional" json:"opportunisticEncryption" yaml:"opportunisticEncryption"`
	// Override the IP/TCP destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#origin Ruleset#origin}
	Origin *RulesetRulesActionParametersOrigin `field:"optional" json:"origin" yaml:"origin"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#origin_cache_control Ruleset#origin_cache_control}
	OriginCacheControl interface{} `field:"optional" json:"originCacheControl" yaml:"originCacheControl"`
	// Generate Cloudflare error pages from issues sent from the origin server.
	//
	// When on, error pages will trigger for issues from the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#origin_error_page_passthru Ruleset#origin_error_page_passthru}
	OriginErrorPagePassthru interface{} `field:"optional" json:"originErrorPagePassthru" yaml:"originErrorPagePassthru"`
	// A set of overrides to apply to the target ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#overrides Ruleset#overrides}
	Overrides *RulesetRulesActionParametersOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// A phase to skip the execution of. This property is only compatible with products. Available values: "current".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#phase Ruleset#phase}
	Phase *string `field:"optional" json:"phase" yaml:"phase"`
	// A list of phases to skip the execution of. This option is incompatible with the rulesets option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#phases Ruleset#phases}
	Phases *[]*string `field:"optional" json:"phases" yaml:"phases"`
	// Configure the Polish level. Available values: "off", "lossless", "lossy", "webp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#polish Ruleset#polish}
	Polish *string `field:"optional" json:"polish" yaml:"polish"`
	// A list of legacy security products to skip the execution of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#products Ruleset#products}
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
	// The raw response fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#raw_response_fields Ruleset#raw_response_fields}
	RawResponseFields interface{} `field:"optional" json:"rawResponseFields" yaml:"rawResponseFields"`
	// Define a timeout value between two successive read operations to your origin server.
	//
	// Historically, the timeout value between two read options from Cloudflare to an origin server is 100 seconds. If you are attempting to reduce HTTP 524 errors because of timeouts from an origin server, try increasing this timeout value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#read_timeout Ruleset#read_timeout}
	ReadTimeout *float64 `field:"optional" json:"readTimeout" yaml:"readTimeout"`
	// The raw request fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#request_fields Ruleset#request_fields}
	RequestFields interface{} `field:"optional" json:"requestFields" yaml:"requestFields"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag) headers.
	//
	// When off, Cloudflare converts strong ETag headers to weak ETag headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#respect_strong_etags Ruleset#respect_strong_etags}
	RespectStrongEtags interface{} `field:"optional" json:"respectStrongEtags" yaml:"respectStrongEtags"`
	// The response to show when the block is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#response Ruleset#response}
	Response *RulesetRulesActionParametersResponse `field:"optional" json:"response" yaml:"response"`
	// The transformed response fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#response_fields Ruleset#response_fields}
	ResponseFields interface{} `field:"optional" json:"responseFields" yaml:"responseFields"`
	// Turn on or off Rocket Loader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#rocket_loader Ruleset#rocket_loader}
	RocketLoader interface{} `field:"optional" json:"rocketLoader" yaml:"rocketLoader"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of.
	//
	// This option is incompatible with the ruleset option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#rules Ruleset#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the rulesets option. Available values: "current".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#ruleset Ruleset#ruleset}
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// A list of ruleset IDs to skip the execution of.
	//
	// This option is incompatible with the ruleset and phases options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#rulesets Ruleset#rulesets}
	Rulesets *[]*string `field:"optional" json:"rulesets" yaml:"rulesets"`
	// Configure the Security Level. Available values: "off", "essentially_off", "low", "medium", "high", "under_attack".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#security_level Ruleset#security_level}
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// Turn on or off Server Side Excludes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#server_side_excludes Ruleset#server_side_excludes}
	ServerSideExcludes interface{} `field:"optional" json:"serverSideExcludes" yaml:"serverSideExcludes"`
	// Define if Cloudflare should serve stale content while getting the latest content from the origin.
	//
	// If on, Cloudflare will not serve stale content while getting the latest content from the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#serve_stale Ruleset#serve_stale}
	ServeStale *RulesetRulesActionParametersServeStale `field:"optional" json:"serveStale" yaml:"serveStale"`
	// Override the Server Name Indication (SNI).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#sni Ruleset#sni}
	Sni *RulesetRulesActionParametersSni `field:"optional" json:"sni" yaml:"sni"`
	// Configure the SSL level. Available values: "off", "flexible", "full", "strict", "origin_pull".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#ssl Ruleset#ssl}
	Ssl *string `field:"optional" json:"ssl" yaml:"ssl"`
	// The status code to use for the error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Turn on or off Signed Exchanges (SXG).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#sxg Ruleset#sxg}
	Sxg interface{} `field:"optional" json:"sxg" yaml:"sxg"`
	// The transformed request fields to log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#transformed_request_fields Ruleset#transformed_request_fields}
	TransformedRequestFields interface{} `field:"optional" json:"transformedRequestFields" yaml:"transformedRequestFields"`
	// URI to rewrite the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/ruleset#uri Ruleset#uri}
	Uri *RulesetRulesActionParametersUri `field:"optional" json:"uri" yaml:"uri"`
}

