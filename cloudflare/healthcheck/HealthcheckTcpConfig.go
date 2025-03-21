// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcheck


type HealthcheckTcpConfig struct {
	// The TCP connection method to use for the health check. Available values: "connection_established".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/healthcheck#method Healthcheck#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/healthcheck#port Healthcheck#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

