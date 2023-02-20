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
	// See [`default_pool_ids`](#default_pool_ids).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#default_pools LoadBalancer#default_pools}
	DefaultPools *[]*string `field:"optional" json:"defaultPools" yaml:"defaultPools"`
	// See [`fallback_pool_id`](#fallback_pool_id).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#fallback_pool LoadBalancer#fallback_pool}
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
	// See [`session_affinity`](#session_affinity).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity LoadBalancer#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// See [`session_affinity_attributes`](#nested-schema-for-session_affinity_attributes). Note that the property [`drain_duration`](#drain_duration) is not currently supported as a rule override.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_attributes LoadBalancer#session_affinity_attributes}
	SessionAffinityAttributes *map[string]*string `field:"optional" json:"sessionAffinityAttributes" yaml:"sessionAffinityAttributes"`
	// See [`session_affinity_ttl`](#session_affinity_ttl).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#session_affinity_ttl LoadBalancer#session_affinity_ttl}
	SessionAffinityTtl *float64 `field:"optional" json:"sessionAffinityTtl" yaml:"sessionAffinityTtl"`
	// See [`steering_policy`](#steering_policy).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#steering_policy LoadBalancer#steering_policy}
	SteeringPolicy *string `field:"optional" json:"steeringPolicy" yaml:"steeringPolicy"`
	// See [`ttl`](#ttl).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/load_balancer#ttl LoadBalancer#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

