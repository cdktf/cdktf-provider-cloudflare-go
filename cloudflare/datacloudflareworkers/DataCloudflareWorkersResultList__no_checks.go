// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datacloudflareworkers

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataCloudflareWorkersResultList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareWorkersResultList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataCloudflareWorkersResultList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareWorkersResultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareWorkersResultList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataCloudflareWorkersResultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataCloudflareWorkersResultListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

