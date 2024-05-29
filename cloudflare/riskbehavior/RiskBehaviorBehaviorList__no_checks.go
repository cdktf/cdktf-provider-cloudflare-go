// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package riskbehavior

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RiskBehaviorBehaviorList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RiskBehaviorBehaviorList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RiskBehaviorBehaviorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RiskBehaviorBehaviorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RiskBehaviorBehaviorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RiskBehaviorBehaviorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RiskBehaviorBehaviorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRiskBehaviorBehaviorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

