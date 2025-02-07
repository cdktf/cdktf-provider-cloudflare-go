// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ratelimit

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RateLimitBypassList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RateLimitBypassList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RateLimitBypassList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RateLimitBypassList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RateLimitBypassList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RateLimitBypassList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRateLimitBypassListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

