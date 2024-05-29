// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessPolicyRequireExternalEvaluationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessPolicyRequireExternalEvaluationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

