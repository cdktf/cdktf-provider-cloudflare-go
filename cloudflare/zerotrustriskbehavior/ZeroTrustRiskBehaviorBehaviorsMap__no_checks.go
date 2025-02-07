// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package zerotrustriskbehavior

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateGetParameters(key *string) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustRiskBehaviorBehaviorsMap) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func validateNewZeroTrustRiskBehaviorBehaviorsMapParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

