// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflarelist

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareListItemsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareListItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareListItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareListItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareListItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

