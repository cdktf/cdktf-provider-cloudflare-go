//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupExcludeSamlList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupExcludeSamlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeSamlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeSamlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeSamlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupExcludeSamlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupExcludeSamlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

