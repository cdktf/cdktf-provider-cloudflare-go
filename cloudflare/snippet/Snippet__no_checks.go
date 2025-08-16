// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package snippet

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Snippet) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Snippet) validatePutFilesParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Snippet) validatePutMetadataParameters(value *SnippetMetadata) error {
	return nil
}

func validateSnippet_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSnippet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSnippet_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSnippet_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetSnippetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Snippet) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewSnippetParameters(scope constructs.Construct, id *string, config *SnippetConfig) error {
	return nil
}

