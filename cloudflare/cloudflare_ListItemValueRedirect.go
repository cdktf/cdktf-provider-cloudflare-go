// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare


type ListItemValueRedirect struct {
	// The source url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#source_url List#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// The target url of the redirect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#target_url List#target_url}
	TargetUrl *string `field:"required" json:"targetUrl" yaml:"targetUrl"`
	// Whether the redirect also matches subdomains of the source url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#include_subdomains List#include_subdomains}
	IncludeSubdomains *string `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Whether to preserve the path suffix when doing subpath matching. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#preserve_path_suffix List#preserve_path_suffix}
	PreservePathSuffix *string `field:"optional" json:"preservePathSuffix" yaml:"preservePathSuffix"`
	// Whether the redirect target url should keep the query string of the request's url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#preserve_query_string List#preserve_query_string}
	PreserveQueryString *string `field:"optional" json:"preserveQueryString" yaml:"preserveQueryString"`
	// The status code to be used when redirecting a request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#status_code List#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Whether the redirect also matches subpaths of the source url. Available values: `disabled`, `enabled`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/list#subpath_matching List#subpath_matching}
	SubpathMatching *string `field:"optional" json:"subpathMatching" yaml:"subpathMatching"`
}

