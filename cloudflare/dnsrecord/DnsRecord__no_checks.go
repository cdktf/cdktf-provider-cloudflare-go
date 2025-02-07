// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnsrecord

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsRecord) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validatePutDataParameters(value *DnsRecordData) error {
	return nil
}

func (d *jsiiProxy_DnsRecord) validatePutSettingsParameters(value *DnsRecordSettings) error {
	return nil
}

func validateDnsRecord_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDnsRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDnsRecord_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDnsRecord_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetContentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetPriorityParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetProxiedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetTtlParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecord) validateSetZoneIdParameters(val *string) error {
	return nil
}

func validateNewDnsRecordParameters(scope constructs.Construct, id *string, config *DnsRecordConfig) error {
	return nil
}

