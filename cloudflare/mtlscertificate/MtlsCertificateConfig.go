// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mtlscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MtlsCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/mtls_certificate#account_id MtlsCertificate#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Indicates whether the certificate is a CA or leaf certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/mtls_certificate#ca MtlsCertificate#ca}
	Ca interface{} `field:"required" json:"ca" yaml:"ca"`
	// The uploaded root CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/mtls_certificate#certificates MtlsCertificate#certificates}
	Certificates *string `field:"required" json:"certificates" yaml:"certificates"`
	// Optional unique name for the certificate. Only used for human readability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/mtls_certificate#name MtlsCertificate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The private key for the certificate.
	//
	// This field is only needed for specific use cases such as using a custom certificate with Zero Trust's block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/mtls_certificate#private_key MtlsCertificate#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
}

