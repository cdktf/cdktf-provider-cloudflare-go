// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptMigrationsSteps struct {
	// A list of classes to delete Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#deleted_classes WorkersScript#deleted_classes}
	DeletedClasses *[]*string `field:"optional" json:"deletedClasses" yaml:"deletedClasses"`
	// A list of classes to create Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#new_classes WorkersScript#new_classes}
	NewClasses *[]*string `field:"optional" json:"newClasses" yaml:"newClasses"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#new_sqlite_classes WorkersScript#new_sqlite_classes}
	NewSqliteClasses *[]*string `field:"optional" json:"newSqliteClasses" yaml:"newSqliteClasses"`
	// A list of classes with Durable Object namespaces that were renamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#renamed_classes WorkersScript#renamed_classes}
	RenamedClasses interface{} `field:"optional" json:"renamedClasses" yaml:"renamedClasses"`
	// A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/workers_script#transferred_classes WorkersScript#transferred_classes}
	TransferredClasses interface{} `field:"optional" json:"transferredClasses" yaml:"transferredClasses"`
}

