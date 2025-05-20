// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package magictransitsiteacl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MagicTransitSiteAclConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#account_id MagicTransitSiteAcl#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#lan_1 MagicTransitSiteAcl#lan_1}.
	Lan1 *MagicTransitSiteAclLan1 `field:"required" json:"lan1" yaml:"lan1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#lan_2 MagicTransitSiteAcl#lan_2}.
	Lan2 *MagicTransitSiteAclLan2 `field:"required" json:"lan2" yaml:"lan2"`
	// The name of the ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#name MagicTransitSiteAcl#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#site_id MagicTransitSiteAcl#site_id}
	SiteId *string `field:"required" json:"siteId" yaml:"siteId"`
	// Description for the ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#description MagicTransitSiteAcl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The desired forwarding action for this ACL policy.
	//
	// If set to "false", the policy will forward traffic to Cloudflare. If set to "true", the policy will forward traffic locally on the Magic Connector. If not included in request, will default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#forward_locally MagicTransitSiteAcl#forward_locally}
	ForwardLocally interface{} `field:"optional" json:"forwardLocally" yaml:"forwardLocally"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#protocols MagicTransitSiteAcl#protocols}.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The desired traffic direction for this ACL policy.
	//
	// If set to "false", the policy will allow bidirectional traffic. If set to "true", the policy will only allow traffic in one direction. If not included in request, will default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.5.0/docs/resources/magic_transit_site_acl#unidirectional MagicTransitSiteAcl#unidirectional}
	Unidirectional interface{} `field:"optional" json:"unidirectional" yaml:"unidirectional"`
}

