// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrusttunnelcloudflaredconfig


type ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#access ZeroTrustTunnelCloudflaredConfigA#access}
	Access *ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccess `field:"optional" json:"access" yaml:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin.
	//
	// This option should be used only if your certificate is not signed by Cloudflare.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#ca_pool ZeroTrustTunnelCloudflaredConfigA#ca_pool}
	CaPool *string `field:"optional" json:"caPool" yaml:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server.
	//
	// This excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#connect_timeout ZeroTrustTunnelCloudflaredConfigA#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#disable_chunked_encoding ZeroTrustTunnelCloudflaredConfigA#disable_chunked_encoding}
	DisableChunkedEncoding interface{} `field:"optional" json:"disableChunkedEncoding" yaml:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#http2_origin ZeroTrustTunnelCloudflaredConfigA#http2_origin}
	Http2Origin interface{} `field:"optional" json:"http2Origin" yaml:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#http_host_header ZeroTrustTunnelCloudflaredConfigA#http_host_header}
	HttpHostHeader *string `field:"optional" json:"httpHostHeader" yaml:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	//
	// This does not restrict the total number of concurrent connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#keep_alive_connections ZeroTrustTunnelCloudflaredConfigA#keep_alive_connections}
	KeepAliveConnections *float64 `field:"optional" json:"keepAliveConnections" yaml:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#keep_alive_timeout ZeroTrustTunnelCloudflaredConfigA#keep_alive_timeout}
	KeepAliveTimeout *float64 `field:"optional" json:"keepAliveTimeout" yaml:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#no_happy_eyeballs ZeroTrustTunnelCloudflaredConfigA#no_happy_eyeballs}
	NoHappyEyeballs interface{} `field:"optional" json:"noHappyEyeballs" yaml:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin.
	//
	// Will allow any certificate from the origin to be accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#no_tls_verify ZeroTrustTunnelCloudflaredConfigA#no_tls_verify}
	NoTlsVerify interface{} `field:"optional" json:"noTlsVerify" yaml:"noTlsVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#origin_server_name ZeroTrustTunnelCloudflaredConfigA#origin_server_name}
	OriginServerName *string `field:"optional" json:"originServerName" yaml:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP.
	//
	// This configures what type of proxy will be started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5 proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#proxy_type ZeroTrustTunnelCloudflaredConfigA#proxy_type}
	ProxyType *string `field:"optional" json:"proxyType" yaml:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#tcp_keep_alive ZeroTrustTunnelCloudflaredConfigA#tcp_keep_alive}
	TcpKeepAlive *float64 `field:"optional" json:"tcpKeepAlive" yaml:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.8.2/docs/resources/zero_trust_tunnel_cloudflared_config#tls_timeout ZeroTrustTunnelCloudflaredConfigA#tls_timeout}
	TlsTimeout *float64 `field:"optional" json:"tlsTimeout" yaml:"tlsTimeout"`
}

