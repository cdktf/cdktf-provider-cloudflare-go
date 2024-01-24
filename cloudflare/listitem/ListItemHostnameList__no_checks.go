// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package listitem

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListItemHostnameList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_ListItemHostnameList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_ListItemHostnameList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ListItemHostnameList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ListItemHostnameList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ListItemHostnameList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ListItemHostnameList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewListItemHostnameListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

