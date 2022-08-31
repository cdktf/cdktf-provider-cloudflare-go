//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyIncludeList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyIncludeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyIncludeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyIncludeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

