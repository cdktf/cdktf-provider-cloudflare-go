// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package zerotrustaccessgroup

// Building without runtime type checking enabled, so all the below just return nil

func (z *jsiiProxy_ZeroTrustAccessGroupRequireList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireList) validateGetParameters(index *float64) error {
	return nil
}

func (z *jsiiProxy_ZeroTrustAccessGroupRequireList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ZeroTrustAccessGroupRequireList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewZeroTrustAccessGroupRequireListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
