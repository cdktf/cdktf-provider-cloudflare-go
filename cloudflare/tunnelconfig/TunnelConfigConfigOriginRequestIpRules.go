package tunnelconfig


type TunnelConfigConfigOriginRequestIpRules struct {
	// Whether to allow the IP prefix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#allow TunnelConfig#allow}
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// Ports to use within the IP rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#ports TunnelConfig#ports}
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// IP rule prefix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/tunnel_config#prefix TunnelConfig#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

