// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustgatewaycertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustGatewayCertificateConfig struct {
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
	// The account identifier to target for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#account_id ZeroTrustGatewayCertificate#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Whether or not to activate a certificate.
	//
	// A certificate must be activated to use in Gateway certificate settings. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#activate ZeroTrustGatewayCertificate#activate}
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// The type of certificate (custom or Gateway-managed). Must provide only one of `custom`, `gateway_managed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#custom ZeroTrustGatewayCertificate#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// The type of certificate (custom or Gateway-managed). Must provide only one of `custom`, `gateway_managed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#gateway_managed ZeroTrustGatewayCertificate#gateway_managed}
	GatewayManaged interface{} `field:"optional" json:"gatewayManaged" yaml:"gatewayManaged"`
	// Certificate UUID. Computed for Gateway-managed certificates. Required when using `custom`. Conflicts with `gateway_managed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#id ZeroTrustGatewayCertificate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Number of days the generated certificate will be valid, minimum 1 day and maximum 30 years.
	//
	// Defaults to 5 years. Defaults to `1826`. Required when using `gateway_managed`. Conflicts with `custom`. **Modifying this attribute will force creation of a new resource.**
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.42.0/docs/resources/zero_trust_gateway_certificate#validity_period_days ZeroTrustGatewayCertificate#validity_period_days}
	ValidityPeriodDays *float64 `field:"optional" json:"validityPeriodDays" yaml:"validityPeriodDays"`
}

