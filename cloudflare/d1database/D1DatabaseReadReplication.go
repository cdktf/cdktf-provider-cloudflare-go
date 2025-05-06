// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package d1database


type D1DatabaseReadReplication struct {
	// The read replication mode for the database.
	//
	// Use 'auto' to create replicas and allow D1 automatically place them around the world, or 'disabled' to not use any database replicas (it can take a few hours for all replicas to be deleted).
	// Available values: "auto", "disabled".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/d1_database#mode D1Database#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

