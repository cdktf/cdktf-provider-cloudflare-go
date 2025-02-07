// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stream

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stream) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Stream) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stream) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateStream_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStream_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStream_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStream_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetAccountIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetAllowedOriginsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetCreatorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetIdentifierParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetMaxDurationSecondsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetMetaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetRequireSignedUrlsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetScheduledDeletionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetThumbnailTimestampPctParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Stream) validateSetUploadExpiryParameters(val *string) error {
	return nil
}

func validateNewStreamParameters(scope constructs.Construct, id *string, config *StreamConfig) error {
	return nil
}

