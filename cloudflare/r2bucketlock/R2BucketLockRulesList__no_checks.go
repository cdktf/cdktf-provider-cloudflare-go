// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package r2bucketlock

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_R2BucketLockRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_R2BucketLockRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_R2BucketLockRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_R2BucketLockRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewR2BucketLockRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

