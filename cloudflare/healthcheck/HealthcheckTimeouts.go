// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcheck


type HealthcheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.37.0/docs/resources/healthcheck#create Healthcheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

