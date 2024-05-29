// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupRequireExternalEvaluationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupRequireExternalEvaluationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

