// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package teamslist

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamsListItemsWithDescriptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TeamsListItemsWithDescriptionList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamsListItemsWithDescriptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamsListItemsWithDescriptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamsListItemsWithDescriptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamsListItemsWithDescriptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamsListItemsWithDescriptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamsListItemsWithDescriptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

