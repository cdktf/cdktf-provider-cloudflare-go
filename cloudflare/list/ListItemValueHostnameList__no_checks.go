// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package list

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListItemValueHostnameList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_ListItemValueHostnameList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_ListItemValueHostnameList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ListItemValueHostnameList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ListItemValueHostnameList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ListItemValueHostnameList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ListItemValueHostnameList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewListItemValueHostnameListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

