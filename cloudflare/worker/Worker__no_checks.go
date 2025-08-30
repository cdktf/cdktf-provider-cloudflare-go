// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Worker) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateImportFromParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (w *jsiiProxy_Worker) validateMoveToIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (w *jsiiProxy_Worker) validatePutObservabilityParameters(value *WorkerObservability) error {
	return nil
}

func (w *jsiiProxy_Worker) validatePutSubdomainParameters(value *WorkerSubdomain) error {
	return nil
}

func (w *jsiiProxy_Worker) validatePutTailConsumersParameters(value interface{}) error {
	return nil
}

func validateWorker_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWorker_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorker_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWorker_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetLogpushParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Worker) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func validateNewWorkerParameters(scope constructs.Construct, id *string, config *WorkerConfig) error {
	return nil
}

