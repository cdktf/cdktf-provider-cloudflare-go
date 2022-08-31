//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadBalancerMonitorHeaderList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadBalancerMonitorHeaderList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerMonitorHeaderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerMonitorHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerMonitorHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadBalancerMonitorHeaderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadBalancerMonitorHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

