// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type RulesetRulesActionParameters struct {
	// browser_ttl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#browser_ttl Ruleset#browser_ttl}
	BrowserTtl *RulesetRulesActionParametersBrowserTtl `field:"optional" json:"browserTtl" yaml:"browserTtl"`
	// Whether to cache if expression matches.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cache Ruleset#cache}
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// cache_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cache_key Ruleset#cache_key}
	CacheKey *RulesetRulesActionParametersCacheKey `field:"optional" json:"cacheKey" yaml:"cacheKey"`
	// Content of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#content Ruleset#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content-Type of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#content_type Ruleset#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// List of cookie values to include as part of custom fields logging.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#cookie_fields Ruleset#cookie_fields}
	CookieFields *[]*string `field:"optional" json:"cookieFields" yaml:"cookieFields"`
	// edge_ttl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#edge_ttl Ruleset#edge_ttl}
	EdgeTtl *RulesetRulesActionParametersEdgeTtl `field:"optional" json:"edgeTtl" yaml:"edgeTtl"`
	// from_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#from_list Ruleset#from_list}
	FromList *RulesetRulesActionParametersFromList `field:"optional" json:"fromList" yaml:"fromList"`
	// from_value block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#from_value Ruleset#from_value}
	FromValue *RulesetRulesActionParametersFromValue `field:"optional" json:"fromValue" yaml:"fromValue"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#headers Ruleset#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Host Header that request origin receives.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#host_header Ruleset#host_header}
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Identifier of the action parameter to modify.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#id Ruleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#increment Ruleset#increment}.
	Increment *float64 `field:"optional" json:"increment" yaml:"increment"`
	// matched_data block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#matched_data Ruleset#matched_data}
	MatchedData *RulesetRulesActionParametersMatchedData `field:"optional" json:"matchedData" yaml:"matchedData"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#origin Ruleset#origin}
	Origin *RulesetRulesActionParametersOrigin `field:"optional" json:"origin" yaml:"origin"`
	// Pass-through error page for origin.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#origin_error_page_passthru Ruleset#origin_error_page_passthru}
	OriginErrorPagePassthru interface{} `field:"optional" json:"originErrorPagePassthru" yaml:"originErrorPagePassthru"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#overrides Ruleset#overrides}
	Overrides *RulesetRulesActionParametersOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// Point in the request/response lifecycle where the ruleset will be created.
	//
	// Available values: `ddos_l4`, `ddos_l7`, `http_custom_errors`, `http_log_custom_fields`, `http_request_cache_settings`, `http_request_firewall_custom`, `http_request_firewall_managed`, `http_request_late_transform`, `http_request_late_transform_managed`, `http_request_main`, `http_request_origin`, `http_request_dynamic_redirect`, `http_request_redirect`, `http_request_sanitize`, `http_request_transform`, `http_response_firewall_managed`, `http_response_headers_transform`, `http_response_headers_transform_managed`, `magic_transit`, `http_ratelimit`, `http_request_sbfm`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#phases Ruleset#phases}
	Phases *[]*string `field:"optional" json:"phases" yaml:"phases"`
	// Products to target with the actions. Available values: `bic`, `hot`, `ratelimit`, `securityLevel`, `uablock`, `waf`, `zonelockdown`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#products Ruleset#products}
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
	// List of request headers to include as part of custom fields logging, in lowercase.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#request_fields Ruleset#request_fields}
	RequestFields *[]*string `field:"optional" json:"requestFields" yaml:"requestFields"`
	// Respect strong ETags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#respect_strong_etags Ruleset#respect_strong_etags}
	RespectStrongEtags interface{} `field:"optional" json:"respectStrongEtags" yaml:"respectStrongEtags"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#response Ruleset#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
	// List of response headers to include as part of custom fields logging, in lowercase.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#response_fields Ruleset#response_fields}
	ResponseFields *[]*string `field:"optional" json:"responseFields" yaml:"responseFields"`
	// Map of managed WAF rule ID to comma-delimited string of ruleset rule IDs.
	//
	// Example: `rules = { "efb7b8c949ac4650a09736fc376e9aee" = "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760" }`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rules Ruleset#rules}
	Rules *map[string]*string `field:"optional" json:"rules" yaml:"rules"`
	// Which ruleset ID to target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#ruleset Ruleset#ruleset}
	Ruleset *string `field:"optional" json:"ruleset" yaml:"ruleset"`
	// List of managed WAF rule IDs to target. Only valid when the `"action"` is set to skip.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#rulesets Ruleset#rulesets}
	Rulesets *[]*string `field:"optional" json:"rulesets" yaml:"rulesets"`
	// serve_stale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#serve_stale Ruleset#serve_stale}
	ServeStale *RulesetRulesActionParametersServeStale `field:"optional" json:"serveStale" yaml:"serveStale"`
	// sni block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#sni Ruleset#sni}
	Sni *RulesetRulesActionParametersSni `field:"optional" json:"sni" yaml:"sni"`
	// HTTP status code of the custom error response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#status_code Ruleset#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// uri block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#uri Ruleset#uri}
	Uri *RulesetRulesActionParametersUri `field:"optional" json:"uri" yaml:"uri"`
	// Version of the ruleset to deploy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/ruleset#version Ruleset#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

