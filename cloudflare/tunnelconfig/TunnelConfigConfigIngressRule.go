package tunnelconfig


type TunnelConfigConfigIngressRule struct {
	// Name of the service to which the request will be sent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#service TunnelConfig#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Hostname to match the incoming request with. If the hostname matches, the request will be sent to the service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#hostname TunnelConfig#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Path of the incoming request. If the path matches, the request will be sent to the local service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#path TunnelConfig#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

