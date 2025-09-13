// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package d1database

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type D1DatabaseConfig struct {
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
	// Account identifier tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/d1_database#account_id D1Database#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// D1 database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/d1_database#name D1Database#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specify the region to create the D1 primary, if available.
	//
	// If this option is omitted, the D1 will be created as close as possible to the current user.
	// Available values: "wnam", "enam", "weur", "eeur", "apac", "oc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/d1_database#primary_location_hint D1Database#primary_location_hint}
	PrimaryLocationHint *string `field:"optional" json:"primaryLocationHint" yaml:"primaryLocationHint"`
	// Configuration for D1 read replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/d1_database#read_replication D1Database#read_replication}
	ReadReplication *D1DatabaseReadReplication `field:"optional" json:"readReplication" yaml:"readReplication"`
}

