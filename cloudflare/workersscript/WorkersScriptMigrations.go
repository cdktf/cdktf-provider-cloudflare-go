// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workersscript


type WorkersScriptMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#deleted_classes WorkersScript#deleted_classes}
	DeletedClasses *[]*string `field:"optional" json:"deletedClasses" yaml:"deletedClasses"`
	// A list of classes to create Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#new_classes WorkersScript#new_classes}
	NewClasses *[]*string `field:"optional" json:"newClasses" yaml:"newClasses"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#new_sqlite_classes WorkersScript#new_sqlite_classes}
	NewSqliteClasses *[]*string `field:"optional" json:"newSqliteClasses" yaml:"newSqliteClasses"`
	// Tag to set as the latest migration tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#new_tag WorkersScript#new_tag}
	NewTag *string `field:"optional" json:"newTag" yaml:"newTag"`
	// Tag used to verify against the latest migration tag for this Worker.
	//
	// If they don't match, the upload is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#old_tag WorkersScript#old_tag}
	OldTag *string `field:"optional" json:"oldTag" yaml:"oldTag"`
	// A list of classes with Durable Object namespaces that were renamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#renamed_classes WorkersScript#renamed_classes}
	RenamedClasses interface{} `field:"optional" json:"renamedClasses" yaml:"renamedClasses"`
	// Migrations to apply in order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#steps WorkersScript#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#transferred_classes WorkersScript#transferred_classes}
	TransferredClasses interface{} `field:"optional" json:"transferredClasses" yaml:"transferredClasses"`
}

