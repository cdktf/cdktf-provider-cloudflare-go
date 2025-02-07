// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedtransforms

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedResponseHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedTransformsManagedResponseHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

