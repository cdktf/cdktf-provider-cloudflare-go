// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamliveinput


type StreamLiveInputRecording struct {
	// Lists the origins allowed to display videos created with this input.
	//
	// Enter allowed origin domains in an array and use `*` for wildcard subdomains. An empty array allows videos to be viewed on any origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/stream_live_input#allowed_origins StreamLiveInput#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Disables reporting the number of live viewers when this property is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/stream_live_input#hide_live_viewer_count StreamLiveInput#hide_live_viewer_count}
	HideLiveViewerCount interface{} `field:"optional" json:"hideLiveViewerCount" yaml:"hideLiveViewerCount"`
	// Specifies the recording behavior for the live input.
	//
	// Set this value to `off` to prevent a recording. Set the value to `automatic` to begin a recording and transition to on-demand after Stream Live stops receiving input.
	// Available values: "off", "automatic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/stream_live_input#mode StreamLiveInput#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Indicates if a video using the live input has the `requireSignedURLs` property set.
	//
	// Also enforces access controls on any video recording of the livestream with the live input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/stream_live_input#require_signed_urls StreamLiveInput#require_signed_urls}
	RequireSignedUrls interface{} `field:"optional" json:"requireSignedUrls" yaml:"requireSignedUrls"`
	// Determines the amount of time a live input configured in `automatic` mode should wait before a recording transitions from live to on-demand.
	//
	// `0` is recommended for most use cases and indicates the platform default should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/stream_live_input#timeout_seconds StreamLiveInput#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

