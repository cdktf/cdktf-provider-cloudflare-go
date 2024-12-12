// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accessapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessApplicationDestinationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationDestinationsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessApplicationDestinationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationDestinationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationDestinationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationDestinationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessApplicationDestinationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessApplicationDestinationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

