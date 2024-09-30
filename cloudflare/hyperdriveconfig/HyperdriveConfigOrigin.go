// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hyperdriveconfig


type HyperdriveConfigOrigin struct {
	// The name of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#database HyperdriveConfig#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The host (hostname or IP) of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#host HyperdriveConfig#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The password of the Hyperdrive configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#password HyperdriveConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the URL scheme used to connect to your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#scheme HyperdriveConfig#scheme}
	Scheme *string `field:"required" json:"scheme" yaml:"scheme"`
	// The user of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#user HyperdriveConfig#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// Client ID associated with the Cloudflare Access Service Token used to connect via Access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#access_client_id HyperdriveConfig#access_client_id}
	AccessClientId *string `field:"optional" json:"accessClientId" yaml:"accessClientId"`
	// Client Secret associated with the Cloudflare Access Service Token used to connect via Access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#access_client_secret HyperdriveConfig#access_client_secret}
	AccessClientSecret *string `field:"optional" json:"accessClientSecret" yaml:"accessClientSecret"`
	// The port (default: 5432 for Postgres) of your origin database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.43.0/docs/resources/hyperdrive_config#port HyperdriveConfig#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

