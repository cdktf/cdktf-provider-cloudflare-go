// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/jsii"

	"github.com/cdktf/cdktf-provider-cloudflare-go/cloudflare/v11/zerotrustdeviceposturerule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDevicePostureRuleInputOutputReference interface {
	cdktf.ComplexObject
	ActiveThreats() *float64
	SetActiveThreats(val *float64)
	ActiveThreatsInput() *float64
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	CheckDisks() *[]*string
	SetCheckDisks(val *[]*string)
	CheckDisksInput() *[]*string
	CheckPrivateKey() interface{}
	SetCheckPrivateKey(val interface{})
	CheckPrivateKeyInput() interface{}
	Cn() *string
	SetCn(val *string)
	CnInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComplianceStatus() *string
	SetComplianceStatus(val *string)
	ComplianceStatusInput() *string
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	CountOperator() *string
	SetCountOperator(val *string)
	CountOperatorInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	EidLastSeen() *string
	SetEidLastSeen(val *string)
	EidLastSeenInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Exists() interface{}
	SetExists(val interface{})
	ExistsInput() interface{}
	ExtendedKeyUsage() *[]*string
	SetExtendedKeyUsage(val *[]*string)
	ExtendedKeyUsageInput() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Infected() interface{}
	SetInfected(val interface{})
	InfectedInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsActive() interface{}
	SetIsActive(val interface{})
	IsActiveInput() interface{}
	IssueCount() *string
	SetIssueCount(val *string)
	IssueCountInput() *string
	LastSeen() *string
	SetLastSeen(val *string)
	LastSeenInput() *string
	Locations() ZeroTrustDevicePostureRuleInputLocationsList
	LocationsInput() interface{}
	NetworkStatus() *string
	SetNetworkStatus(val *string)
	NetworkStatusInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Os() *string
	SetOs(val *string)
	OsDistroName() *string
	SetOsDistroName(val *string)
	OsDistroNameInput() *string
	OsDistroRevision() *string
	SetOsDistroRevision(val *string)
	OsDistroRevisionInput() *string
	OsInput() *string
	OsVersionExtra() *string
	SetOsVersionExtra(val *string)
	OsVersionExtraInput() *string
	Overall() *string
	SetOverall(val *string)
	OverallInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RequireAll() interface{}
	SetRequireAll(val interface{})
	RequireAllInput() interface{}
	RiskLevel() *string
	SetRiskLevel(val *string)
	RiskLevelInput() *string
	Running() interface{}
	SetRunning(val interface{})
	RunningInput() interface{}
	Score() *float64
	SetScore(val *float64)
	ScoreInput() *float64
	SensorConfig() *string
	SetSensorConfig(val *string)
	SensorConfigInput() *string
	Sha256() *string
	SetSha256(val *string)
	Sha256Input() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
	TotalScore() *float64
	SetTotalScore(val *float64)
	TotalScoreInput() *float64
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	VersionOperator() *string
	SetVersionOperator(val *string)
	VersionOperatorInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutLocations(value interface{})
	ResetActiveThreats()
	ResetCertificateId()
	ResetCheckDisks()
	ResetCheckPrivateKey()
	ResetCn()
	ResetComplianceStatus()
	ResetConnectionId()
	ResetCountOperator()
	ResetDomain()
	ResetEidLastSeen()
	ResetEnabled()
	ResetExists()
	ResetExtendedKeyUsage()
	ResetId()
	ResetInfected()
	ResetIsActive()
	ResetIssueCount()
	ResetLastSeen()
	ResetLocations()
	ResetNetworkStatus()
	ResetOperator()
	ResetOs()
	ResetOsDistroName()
	ResetOsDistroRevision()
	ResetOsVersionExtra()
	ResetOverall()
	ResetPath()
	ResetRequireAll()
	ResetRiskLevel()
	ResetRunning()
	ResetScore()
	ResetSensorConfig()
	ResetSha256()
	ResetState()
	ResetThumbprint()
	ResetTotalScore()
	ResetVersion()
	ResetVersionOperator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ZeroTrustDevicePostureRuleInputOutputReference
type jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ActiveThreats() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeThreats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ActiveThreatsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeThreatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CheckDisks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CheckDisksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checkDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CheckPrivateKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CheckPrivateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Cn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ComplianceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ComplianceStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CountOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CountOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) EidLastSeen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eidLastSeen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) EidLastSeenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eidLastSeenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Exists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"existsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ExtendedKeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ExtendedKeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Infected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) InfectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) IsActive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) IsActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isActiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) IssueCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) IssueCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) LastSeen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSeen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) LastSeenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastSeenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Locations() ZeroTrustDevicePostureRuleInputLocationsList {
	var returns ZeroTrustDevicePostureRuleInputLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) LocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) NetworkStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) NetworkStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsDistroName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsDistroNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsDistroRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsDistroRevisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDistroRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsVersionExtra() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersionExtra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OsVersionExtraInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersionExtraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Overall() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) OverallInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) RequireAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) RequireAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) RiskLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riskLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) RiskLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"riskLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Running() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"running",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) RunningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Score() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"score",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) SensorConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) SensorConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sensorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Sha256Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) TotalScore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) TotalScoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) VersionOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) VersionOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionOperatorInput",
		&returns,
	)
	return returns
}


func NewZeroTrustDevicePostureRuleInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ZeroTrustDevicePostureRuleInputOutputReference {
	_init_.Initialize()

	if err := validateNewZeroTrustDevicePostureRuleInputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewZeroTrustDevicePostureRuleInputOutputReference_Override(z ZeroTrustDevicePostureRuleInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetActiveThreats(val *float64) {
	if err := j.validateSetActiveThreatsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeThreats",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetCertificateId(val *string) {
	if err := j.validateSetCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetCheckDisks(val *[]*string) {
	if err := j.validateSetCheckDisksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkDisks",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetCheckPrivateKey(val interface{}) {
	if err := j.validateSetCheckPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkPrivateKey",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetCn(val *string) {
	if err := j.validateSetCnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetComplianceStatus(val *string) {
	if err := j.validateSetComplianceStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceStatus",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetCountOperator(val *string) {
	if err := j.validateSetCountOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countOperator",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetEidLastSeen(val *string) {
	if err := j.validateSetEidLastSeenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eidLastSeen",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetExists(val interface{}) {
	if err := j.validateSetExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exists",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetExtendedKeyUsage(val *[]*string) {
	if err := j.validateSetExtendedKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedKeyUsage",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetInfected(val interface{}) {
	if err := j.validateSetInfectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infected",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetIsActive(val interface{}) {
	if err := j.validateSetIsActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isActive",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetIssueCount(val *string) {
	if err := j.validateSetIssueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueCount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetLastSeen(val *string) {
	if err := j.validateSetLastSeenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastSeen",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetNetworkStatus(val *string) {
	if err := j.validateSetNetworkStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkStatus",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOsDistroName(val *string) {
	if err := j.validateSetOsDistroNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDistroName",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOsDistroRevision(val *string) {
	if err := j.validateSetOsDistroRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDistroRevision",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOsVersionExtra(val *string) {
	if err := j.validateSetOsVersionExtraParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osVersionExtra",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetOverall(val *string) {
	if err := j.validateSetOverallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overall",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetRequireAll(val interface{}) {
	if err := j.validateSetRequireAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAll",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetRiskLevel(val *string) {
	if err := j.validateSetRiskLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"riskLevel",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetRunning(val interface{}) {
	if err := j.validateSetRunningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"running",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetScore(val *float64) {
	if err := j.validateSetScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"score",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetSensorConfig(val *string) {
	if err := j.validateSetSensorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensorConfig",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetSha256(val *string) {
	if err := j.validateSetSha256Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sha256",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetThumbprint(val *string) {
	if err := j.validateSetThumbprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetTotalScore(val *float64) {
	if err := j.validateSetTotalScoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalScore",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference)SetVersionOperator(val *string) {
	if err := j.validateSetVersionOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionOperator",
		val,
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) PutLocations(value interface{}) {
	if err := z.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLocations",
		[]interface{}{value},
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetActiveThreats() {
	_jsii_.InvokeVoid(
		z,
		"resetActiveThreats",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetCertificateId() {
	_jsii_.InvokeVoid(
		z,
		"resetCertificateId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetCheckDisks() {
	_jsii_.InvokeVoid(
		z,
		"resetCheckDisks",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetCheckPrivateKey() {
	_jsii_.InvokeVoid(
		z,
		"resetCheckPrivateKey",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetCn() {
	_jsii_.InvokeVoid(
		z,
		"resetCn",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetComplianceStatus() {
	_jsii_.InvokeVoid(
		z,
		"resetComplianceStatus",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		z,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetCountOperator() {
	_jsii_.InvokeVoid(
		z,
		"resetCountOperator",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetEidLastSeen() {
	_jsii_.InvokeVoid(
		z,
		"resetEidLastSeen",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetExists() {
	_jsii_.InvokeVoid(
		z,
		"resetExists",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetExtendedKeyUsage() {
	_jsii_.InvokeVoid(
		z,
		"resetExtendedKeyUsage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		z,
		"resetId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetInfected() {
	_jsii_.InvokeVoid(
		z,
		"resetInfected",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetIsActive() {
	_jsii_.InvokeVoid(
		z,
		"resetIsActive",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetIssueCount() {
	_jsii_.InvokeVoid(
		z,
		"resetIssueCount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetLastSeen() {
	_jsii_.InvokeVoid(
		z,
		"resetLastSeen",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		z,
		"resetLocations",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetNetworkStatus() {
	_jsii_.InvokeVoid(
		z,
		"resetNetworkStatus",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		z,
		"resetOperator",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOs() {
	_jsii_.InvokeVoid(
		z,
		"resetOs",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOsDistroName() {
	_jsii_.InvokeVoid(
		z,
		"resetOsDistroName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOsDistroRevision() {
	_jsii_.InvokeVoid(
		z,
		"resetOsDistroRevision",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOsVersionExtra() {
	_jsii_.InvokeVoid(
		z,
		"resetOsVersionExtra",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetOverall() {
	_jsii_.InvokeVoid(
		z,
		"resetOverall",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		z,
		"resetPath",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetRequireAll() {
	_jsii_.InvokeVoid(
		z,
		"resetRequireAll",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetRiskLevel() {
	_jsii_.InvokeVoid(
		z,
		"resetRiskLevel",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetRunning() {
	_jsii_.InvokeVoid(
		z,
		"resetRunning",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetScore() {
	_jsii_.InvokeVoid(
		z,
		"resetScore",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetSensorConfig() {
	_jsii_.InvokeVoid(
		z,
		"resetSensorConfig",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetSha256() {
	_jsii_.InvokeVoid(
		z,
		"resetSha256",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		z,
		"resetState",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetThumbprint() {
	_jsii_.InvokeVoid(
		z,
		"resetThumbprint",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetTotalScore() {
	_jsii_.InvokeVoid(
		z,
		"resetTotalScore",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		z,
		"resetVersion",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ResetVersionOperator() {
	_jsii_.InvokeVoid(
		z,
		"resetVersionOperator",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := z.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		z,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

