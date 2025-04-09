// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatepack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificatePackConfig struct {
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
	// Certificate Authority selected for the order.
	//
	// For information on any certificate authority specific details or restrictions [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	// Available values: "google", "lets_encrypt", "ssl_com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#certificate_authority CertificatePack#certificate_authority}
	CertificateAuthority *string `field:"required" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Comma separated list of valid host names for the certificate packs.
	//
	// Must contain the zone apex, may not contain more than 50 hosts, and may not be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#hosts CertificatePack#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Type of certificate pack. Available values: "advanced".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#type CertificatePack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Validation Method selected for the order. Available values: "txt", "http", "email".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#validation_method CertificatePack#validation_method}
	ValidationMethod *string `field:"required" json:"validationMethod" yaml:"validationMethod"`
	// Validity Days selected for the order. Available values: 14, 30, 90, 365.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#validity_days CertificatePack#validity_days}
	ValidityDays *float64 `field:"required" json:"validityDays" yaml:"validityDays"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#zone_id CertificatePack#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Whether or not to add Cloudflare Branding for the order.
	//
	// This will add a subdomain of sni.cloudflaressl.com as the Common Name if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.3.0/docs/resources/certificate_pack#cloudflare_branding CertificatePack#cloudflare_branding}
	CloudflareBranding interface{} `field:"optional" json:"cloudflareBranding" yaml:"cloudflareBranding"`
}

