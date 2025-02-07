// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accountmember

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountMemberPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountMemberPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountMemberPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountMemberPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountMemberPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

