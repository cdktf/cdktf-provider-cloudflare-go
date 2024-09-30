// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accessapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessApplicationTargetCriteriaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationTargetCriteriaList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationTargetCriteriaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationTargetCriteriaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationTargetCriteriaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationTargetCriteriaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationTargetCriteriaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessApplicationTargetCriteriaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

