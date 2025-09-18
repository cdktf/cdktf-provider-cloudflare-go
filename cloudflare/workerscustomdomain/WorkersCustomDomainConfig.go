// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerscustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersCustomDomainConfig struct {
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
	// Identifer of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_custom_domain#account_id WorkersCustomDomain#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Worker environment associated with the zone and hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_custom_domain#environment WorkersCustomDomain#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Hostname of the Worker Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_custom_domain#hostname WorkersCustomDomain#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Worker service associated with the zone and hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_custom_domain#service WorkersCustomDomain#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Identifier of the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/workers_custom_domain#zone_id WorkersCustomDomain#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

