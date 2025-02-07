// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ratelimit

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RateLimitMatchHeadersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RateLimitMatchHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RateLimitMatchHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RateLimitMatchHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RateLimitMatchHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RateLimitMatchHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RateLimitMatchHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRateLimitMatchHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

