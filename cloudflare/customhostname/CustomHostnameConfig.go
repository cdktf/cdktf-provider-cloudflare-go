// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customhostname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomHostnameConfig struct {
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
	// The custom hostname that will point to your hostname via CNAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#hostname CustomHostname#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// SSL properties used when creating the custom hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#ssl CustomHostname#ssl}
	Ssl *CustomHostnameSsl `field:"required" json:"ssl" yaml:"ssl"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#zone_id CustomHostname#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer) settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#custom_metadata CustomHostname#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#custom_origin_server CustomHostname#custom_origin_server}
	CustomOriginServer *string `field:"optional" json:"customOriginServer" yaml:"customOriginServer"`
	// A hostname that will be sent to your custom origin server as SNI for TLS handshake.
	//
	// This can be a valid subdomain of the zone or custom origin server name or the string ':request_host_header:' which will cause the host header in the request to be used as SNI. Not configurable with default/fallback origin server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.6.0/docs/resources/custom_hostname#custom_origin_sni CustomHostname#custom_origin_sni}
	CustomOriginSni *string `field:"optional" json:"customOriginSni" yaml:"customOriginSni"`
}

