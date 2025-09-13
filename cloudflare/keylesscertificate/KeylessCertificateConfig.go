// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keylesscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeylessCertificateConfig struct {
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
	// The zone's SSL certificate or SSL certificate and intermediate(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#certificate KeylessCertificate#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The keyless SSL name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#host KeylessCertificate#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#zone_id KeylessCertificate#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores.
	//
	// An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.
	// Available values: "ubiquitous", "optimal", "force".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#bundle_method KeylessCertificate#bundle_method}
	BundleMethod *string `field:"optional" json:"bundleMethod" yaml:"bundleMethod"`
	// Whether or not the Keyless SSL is on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#enabled KeylessCertificate#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The keyless SSL name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#name KeylessCertificate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#port KeylessCertificate#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.10.0/docs/resources/keyless_certificate#tunnel KeylessCertificate#tunnel}
	Tunnel *KeylessCertificateTunnel `field:"optional" json:"tunnel" yaml:"tunnel"`
}

