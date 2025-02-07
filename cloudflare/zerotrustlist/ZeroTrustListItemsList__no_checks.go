// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package zerotrustlist

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZeroTrustListItemsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustListItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustListItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustListItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustListItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustListItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustListItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewZeroTrustListItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

