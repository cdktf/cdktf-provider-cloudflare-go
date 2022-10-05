package loadbalancer


type LoadBalancerRulesOverrides struct {
	// adaptive_routing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#adaptive_routing LoadBalancer#adaptive_routing}
	AdaptiveRouting interface{} `field:"optional" json:"adaptiveRouting" yaml:"adaptiveRouting"`
	// country_pools block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#country_pools LoadBalancer#country_pools}
	CountryPools interface{} `field:"optional" json:"countryPools" yaml:"countryPools"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#default_pools LoadBalancer#default_pools}.
	DefaultPools *[]*string `field:"optional" json:"defaultPools" yaml:"defaultPools"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#fallback_pool LoadBalancer#fallback_pool}.
	FallbackPool *string `field:"optional" json:"fallbackPool" yaml:"fallbackPool"`
	// location_strategy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#location_strategy LoadBalancer#location_strategy}
	LocationStrategy interface{} `field:"optional" json:"locationStrategy" yaml:"locationStrategy"`
	// pop_pools block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#pop_pools LoadBalancer#pop_pools}
	PopPools interface{} `field:"optional" json:"popPools" yaml:"popPools"`
	// random_steering block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#random_steering LoadBalancer#random_steering}
	RandomSteering interface{} `field:"optional" json:"randomSteering" yaml:"randomSteering"`
	// region_pools block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#region_pools LoadBalancer#region_pools}
	RegionPools interface{} `field:"optional" json:"regionPools" yaml:"regionPools"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity LoadBalancer#session_affinity}.
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}.
	SessionAffinityAttributes *map[string]*string `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}.
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#steering_policy LoadBalancer#steering_policy}.
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#ttl LoadBalancer#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

