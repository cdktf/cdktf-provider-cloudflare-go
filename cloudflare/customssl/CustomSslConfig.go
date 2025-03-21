// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customssl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomSslConfig struct {
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
	// The zone's SSL certificate or certificate and the intermediate(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#certificate CustomSsl#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The zone's private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#private_key CustomSsl#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#zone_id CustomSsl#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores.
	//
	// An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.
	// Available values: "ubiquitous", "optimal", "force".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#bundle_method CustomSsl#bundle_method}
	BundleMethod *string `field:"optional" json:"bundleMethod" yaml:"bundleMethod"`
	// Specify the region where your private key can be held locally for optimal TLS performance.
	//
	// HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Options allow distribution to only to U.S. data centers, only to E.U. data centers, or only to highest security data centers. Default distribution is to all Cloudflare datacenters, for optimal performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#geo_restrictions CustomSsl#geo_restrictions}
	GeoRestrictions *CustomSslGeoRestrictions `field:"optional" json:"geoRestrictions" yaml:"geoRestrictions"`
	// Specify the policy that determines the region where your private key will be held locally.
	//
	// HTTPS connections to any excluded data center will still be fully encrypted, but will incur some latency while Keyless SSL is used to complete the handshake with the nearest allowed data center. Any combination of countries, specified by their two letter country code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Officially_assigned_code_elements) can be chosen, such as 'country: IN', as well as 'region: EU' which refers to the EU region. If there are too few data centers satisfying the policy, it will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#policy CustomSsl#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// The type 'legacy_custom' enables support for legacy clients which do not include SNI in the TLS handshake.
	//
	// Available values: "legacy_custom", "sni_custom".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.2.0/docs/resources/custom_ssl#type CustomSsl#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

