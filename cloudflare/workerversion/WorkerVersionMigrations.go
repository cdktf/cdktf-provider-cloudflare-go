// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workerversion


type WorkerVersionMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#deleted_classes WorkerVersion#deleted_classes}
	DeletedClasses *[]*string `field:"optional" json:"deletedClasses" yaml:"deletedClasses"`
	// A list of classes to create Durable Object namespaces from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#new_classes WorkerVersion#new_classes}
	NewClasses *[]*string `field:"optional" json:"newClasses" yaml:"newClasses"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#new_sqlite_classes WorkerVersion#new_sqlite_classes}
	NewSqliteClasses *[]*string `field:"optional" json:"newSqliteClasses" yaml:"newSqliteClasses"`
	// Tag to set as the latest migration tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#new_tag WorkerVersion#new_tag}
	NewTag *string `field:"optional" json:"newTag" yaml:"newTag"`
	// Tag used to verify against the latest migration tag for this Worker.
	//
	// If they don't match, the upload is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#old_tag WorkerVersion#old_tag}
	OldTag *string `field:"optional" json:"oldTag" yaml:"oldTag"`
	// A list of classes with Durable Object namespaces that were renamed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#renamed_classes WorkerVersion#renamed_classes}
	RenamedClasses interface{} `field:"optional" json:"renamedClasses" yaml:"renamedClasses"`
	// Migrations to apply in order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#steps WorkerVersion#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.1/docs/resources/worker_version#transferred_classes WorkerVersion#transferred_classes}
	TransferredClasses interface{} `field:"optional" json:"transferredClasses" yaml:"transferredClasses"`
}

