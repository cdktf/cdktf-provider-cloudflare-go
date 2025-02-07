// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accounttoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesPermissionGroupsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountTokenPoliciesPermissionGroupsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

