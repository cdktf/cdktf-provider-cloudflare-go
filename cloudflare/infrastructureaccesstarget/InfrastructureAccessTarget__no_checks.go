// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package infrastructureaccesstarget

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InfrastructureAccessTarget) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateMoveToIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_InfrastructureAccessTarget) validatePutIpParameters(value *InfrastructureAccessTargetIp) error {
	return nil
}

func validateInfrastructureAccessTarget_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateInfrastructureAccessTarget_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInfrastructureAccessTarget_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateInfrastructureAccessTarget_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetHostnameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_InfrastructureAccessTarget) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewInfrastructureAccessTargetParameters(scope constructs.Construct, id *string, config *InfrastructureAccessTargetConfig) error {
	return nil
}

