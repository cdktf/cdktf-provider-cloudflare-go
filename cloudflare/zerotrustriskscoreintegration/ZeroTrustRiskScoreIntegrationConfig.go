// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustriskscoreintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustRiskScoreIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_risk_score_integration#account_id ZeroTrustRiskScoreIntegration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The type of integration, e.g. 'Okta'. Full list of allowed values can be found here: https://developers.cloudflare.com/api/operations/dlp-zt-risk-score-integration-create#request-body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_risk_score_integration#integration_type ZeroTrustRiskScoreIntegration#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The base url of the tenant, e.g. 'https://tenant.okta.com'. Must be your Okta Tenant URL and not your custom domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_risk_score_integration#tenant_url ZeroTrustRiskScoreIntegration#tenant_url}
	TenantUrl *string `field:"required" json:"tenantUrl" yaml:"tenantUrl"`
	// Whether this integration is enabled. If disabled, no risk changes will be exported to the third-party.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_risk_score_integration#active ZeroTrustRiskScoreIntegration#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// A reference id that can be supplied by the client.
	//
	// Currently this should be set to the Access-Okta IDP ID (a UUIDv4). If omitted, a random UUIDv4 is used. https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.0/docs/resources/zero_trust_risk_score_integration#reference_id ZeroTrustRiskScoreIntegration#reference_id}
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
}

