// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcheckConfig struct {
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
	// The hostname or IP address of the origin server to run health checks on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#address Healthcheck#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#name Healthcheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#zone_id Healthcheck#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A list of regions from which to run health checks. Null means Cloudflare will pick a default region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#check_regions Healthcheck#check_regions}
	CheckRegions *[]*string `field:"optional" json:"checkRegions" yaml:"checkRegions"`
	// The number of consecutive fails required from a health check before changing the health to unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#consecutive_fails Healthcheck#consecutive_fails}
	ConsecutiveFails *float64 `field:"optional" json:"consecutiveFails" yaml:"consecutiveFails"`
	// The number of consecutive successes required from a health check before changing the health to healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#consecutive_successes Healthcheck#consecutive_successes}
	ConsecutiveSuccesses *float64 `field:"optional" json:"consecutiveSuccesses" yaml:"consecutiveSuccesses"`
	// A human-readable description of the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#description Healthcheck#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#http_config Healthcheck#http_config}
	HttpConfig *HealthcheckHttpConfig `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// The interval between each health check.
	//
	// Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#interval Healthcheck#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy.
	//
	// Retries are attempted immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#retries Healthcheck#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// If suspended, no health checks are sent to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#suspended Healthcheck#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// Parameters specific to TCP health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#tcp_config Healthcheck#tcp_config}
	TcpConfig *HealthcheckTcpConfig `field:"optional" json:"tcpConfig" yaml:"tcpConfig"`
	// The timeout (in seconds) before marking the health check as failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#timeout Healthcheck#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/healthcheck#type Healthcheck#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

