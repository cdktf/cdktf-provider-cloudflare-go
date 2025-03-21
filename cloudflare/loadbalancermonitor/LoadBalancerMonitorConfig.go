// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancermonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadBalancerMonitorConfig struct {
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
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#account_id LoadBalancerMonitor#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Do not validate the certificate when monitor use HTTPS.
	//
	// This parameter is currently only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#allow_insecure LoadBalancerMonitor#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N consecutive times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#consecutive_down LoadBalancerMonitor#consecutive_down}
	ConsecutiveDown *float64 `field:"optional" json:"consecutiveDown" yaml:"consecutiveDown"`
	// To be marked healthy the monitored origin must pass this healthcheck N consecutive times.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#consecutive_up LoadBalancerMonitor#consecutive_up}
	ConsecutiveUp *float64 `field:"optional" json:"consecutiveUp" yaml:"consecutiveUp"`
	// Object description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#description LoadBalancerMonitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A case-insensitive sub-string to look for in the response body.
	//
	// If this string is not found, the origin will be marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#expected_body LoadBalancerMonitor#expected_body}
	ExpectedBody *string `field:"optional" json:"expectedBody" yaml:"expectedBody"`
	// The expected HTTP response code or code range of the health check.
	//
	// This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#expected_codes LoadBalancerMonitor#expected_codes}
	ExpectedCodes *string `field:"optional" json:"expectedCodes" yaml:"expectedCodes"`
	// Follow redirects if returned by the origin. This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#follow_redirects LoadBalancerMonitor#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// The HTTP request headers to send in the health check.
	//
	// It is recommended you set a Host header by default. The User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#header LoadBalancerMonitor#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The interval between each health check.
	//
	// Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#interval LoadBalancerMonitor#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The method to use for the health check.
	//
	// This defaults to 'GET' for HTTP/HTTPS based checks and 'connection_established' for TCP based health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#method LoadBalancerMonitor#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The endpoint path you want to conduct a health check against.
	//
	// This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#path LoadBalancerMonitor#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port number to connect to for the health check.
	//
	// Required for TCP, UDP, and SMTP checks. HTTP and HTTPS checks should only define the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#port LoadBalancerMonitor#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Assign this monitor to emulate the specified zone while probing.
	//
	// This parameter is only valid for HTTP and HTTPS monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#probe_zone LoadBalancerMonitor#probe_zone}
	ProbeZone *string `field:"optional" json:"probeZone" yaml:"probeZone"`
	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
	//
	// Retries are attempted immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#retries LoadBalancerMonitor#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#timeout LoadBalancerMonitor#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The protocol to use for the health check.
	//
	// Currently supported protocols are 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	// Available values: "http", "https", "tcp", "udp_icmp", "icmp_ping", "smtp".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/load_balancer_monitor#type LoadBalancerMonitor#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

