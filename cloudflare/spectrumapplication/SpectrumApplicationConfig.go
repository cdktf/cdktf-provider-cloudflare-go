// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spectrumapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpectrumApplicationConfig struct {
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
	// The name and type of DNS record for the Spectrum application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#dns SpectrumApplication#dns}
	Dns *SpectrumApplicationDns `field:"required" json:"dns" yaml:"dns"`
	// The port configuration at Cloudflare's edge.
	//
	// May specify a single port, for example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#protocol SpectrumApplication#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Zone identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#zone_id SpectrumApplication#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP applications with traffic_type set to "direct".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#argo_smart_routing SpectrumApplication#argo_smart_routing}
	ArgoSmartRouting interface{} `field:"optional" json:"argoSmartRouting" yaml:"argoSmartRouting"`
	// The anycast edge IP configuration for the hostname of this application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#edge_ips SpectrumApplication#edge_ips}
	EdgeIps *SpectrumApplicationEdgeIps `field:"optional" json:"edgeIps" yaml:"edgeIps"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP applications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#ip_firewall SpectrumApplication#ip_firewall}
	IpFirewall interface{} `field:"optional" json:"ipFirewall" yaml:"ipFirewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#origin_direct SpectrumApplication#origin_direct}
	OriginDirect *[]*string `field:"optional" json:"originDirect" yaml:"originDirect"`
	// The name and type of DNS record for the Spectrum application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#origin_dns SpectrumApplication#origin_dns}
	OriginDns *SpectrumApplicationOriginDns `field:"optional" json:"originDns" yaml:"originDns"`
	// The destination port at the origin.
	//
	// Only specified in conjunction with origin_dns. May use an integer to specify a single origin port, for example `1000`, or a string to specify a range of origin ports, for example `"1000-2000"`.
	// Notes: If specifying a port range, the number of ports in the range must match the number of ports specified in the "protocol" field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#origin_port SpectrumApplication#origin_port}
	OriginPort *map[string]interface{} `field:"optional" json:"originPort" yaml:"originPort"`
	// Enables Proxy Protocol to the origin.
	//
	// Refer to [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/) for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple Proxy Protocol.
	// Available values: "off", "v1", "v2", "simple".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#proxy_protocol SpectrumApplication#proxy_protocol}
	ProxyProtocol *string `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
	// The type of TLS termination associated with the application. Available values: "off", "flexible", "full", "strict".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#tls SpectrumApplication#tls}
	Tls *string `field:"optional" json:"tls" yaml:"tls"`
	// Determines how data travels from the edge to your origin.
	//
	// When set to "direct", Spectrum will send traffic directly to your origin, and the application's type is derived from the `protocol`. When set to "http" or "https", Spectrum will apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and the application type matches this property exactly.
	// Available values: "direct", "http", "https".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.4.0/docs/resources/spectrum_application#traffic_type SpectrumApplication#traffic_type}
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

