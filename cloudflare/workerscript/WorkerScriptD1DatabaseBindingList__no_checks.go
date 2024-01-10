// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workerscript

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerScriptD1DatabaseBindingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerScriptD1DatabaseBindingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

