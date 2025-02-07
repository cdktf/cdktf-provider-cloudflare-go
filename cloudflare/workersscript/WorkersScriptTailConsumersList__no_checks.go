// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workersscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersScriptTailConsumersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptTailConsumersList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersScriptTailConsumersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptTailConsumersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptTailConsumersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptTailConsumersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersScriptTailConsumersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersScriptTailConsumersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

