//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt cloudflare Provider for Terraform CDK (cdktf)
package cloudflare

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (z *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ZoneSettingsOverrideInitialSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewZoneSettingsOverrideInitialSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

