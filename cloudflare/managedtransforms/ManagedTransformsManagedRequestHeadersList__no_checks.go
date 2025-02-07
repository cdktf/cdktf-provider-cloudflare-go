// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedtransforms

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedTransformsManagedRequestHeadersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedTransformsManagedRequestHeadersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

