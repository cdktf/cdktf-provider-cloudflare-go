// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package queueconsumer

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueConsumer) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateImportFromParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateMoveToIdParameters(id *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (q *jsiiProxy_QueueConsumer) validatePutSettingsParameters(value *QueueConsumerSettings) error {
	return nil
}

func validateQueueConsumer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateQueueConsumer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQueueConsumer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateQueueConsumer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetConsumerIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetDeadLetterQueueParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetQueueIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetScriptNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QueueConsumer) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewQueueConsumerParameters(scope constructs.Construct, id *string, config *QueueConsumerConfig) error {
	return nil
}

