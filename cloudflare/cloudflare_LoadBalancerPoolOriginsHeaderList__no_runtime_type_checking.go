//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerPoolOriginsHeaderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerPoolOriginsHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
