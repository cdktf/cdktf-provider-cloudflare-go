// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessGroupIncludeExternalEvaluationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessGroupIncludeExternalEvaluationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

