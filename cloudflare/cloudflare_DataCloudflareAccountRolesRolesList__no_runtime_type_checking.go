//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareAccountRolesRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareAccountRolesRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareAccountRolesRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareAccountRolesRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareAccountRolesRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareAccountRolesRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

