// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apitoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiTokenPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiTokenPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiTokenPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiTokenPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

