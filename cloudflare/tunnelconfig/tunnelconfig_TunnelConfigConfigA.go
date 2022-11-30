package tunnelconfig


type TunnelConfigConfigA struct {
	// ingress_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#ingress_rule TunnelConfig#ingress_rule}
	IngressRule interface{} `field:"required" json:"ingressRule" yaml:"ingressRule"`
	// origin_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#origin_request TunnelConfig#origin_request}
	OriginRequest *TunnelConfigConfigOriginRequest `field:"optional" json:"originRequest" yaml:"originRequest"`
	// warp_routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#warp_routing TunnelConfig#warp_routing}
	WarpRouting *TunnelConfigConfigWarpRouting `field:"optional" json:"warpRouting" yaml:"warpRouting"`
}

