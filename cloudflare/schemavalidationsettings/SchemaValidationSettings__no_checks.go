// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schemavalidationsettings

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaValidationSettings) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SchemaValidationSettings) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateSchemaValidationSettings_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSchemaValidationSettings_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSchemaValidationSettings_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSchemaValidationSettings_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetValidationDefaultMitigationActionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetValidationOverrideMitigationActionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaValidationSettings) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewSchemaValidationSettingsParameters(scope constructs.Construct, id *string, config *SchemaValidationSettingsConfig) error {
	return nil
}

