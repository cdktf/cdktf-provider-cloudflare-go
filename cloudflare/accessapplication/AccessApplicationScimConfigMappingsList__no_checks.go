// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accessapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessApplicationScimConfigMappingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationScimConfigMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationScimConfigMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationScimConfigMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationScimConfigMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationScimConfigMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationScimConfigMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessApplicationScimConfigMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

