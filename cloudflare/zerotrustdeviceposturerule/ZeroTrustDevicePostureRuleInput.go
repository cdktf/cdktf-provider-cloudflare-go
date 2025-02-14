// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleInput struct {
	// The Number of active threats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#active_threats ZeroTrustDevicePostureRule#active_threats}
	ActiveThreats *float64 `field:"optional" json:"activeThreats" yaml:"activeThreats"`
	// UUID of Cloudflare managed certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#certificate_id ZeroTrustDevicePostureRule#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// List of volume names to be checked for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#check_disks ZeroTrustDevicePostureRule#check_disks}
	CheckDisks *[]*string `field:"optional" json:"checkDisks" yaml:"checkDisks"`
	// Confirm the certificate was not imported from another device.
	//
	// We recommend keeping this enabled unless the certificate was deployed without a private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#check_private_key ZeroTrustDevicePostureRule#check_private_key}
	CheckPrivateKey interface{} `field:"optional" json:"checkPrivateKey" yaml:"checkPrivateKey"`
	// Common Name that is protected by the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#cn ZeroTrustDevicePostureRule#cn}
	Cn *string `field:"optional" json:"cn" yaml:"cn"`
	// Compliance Status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#compliance_status ZeroTrustDevicePostureRule#compliance_status}
	ComplianceStatus *string `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// Posture Integration ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#connection_id ZeroTrustDevicePostureRule#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// Count Operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#count_operator ZeroTrustDevicePostureRule#count_operator}
	CountOperator *string `field:"optional" json:"countOperator" yaml:"countOperator"`
	// Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#domain ZeroTrustDevicePostureRule#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// For more details on eid last seen, refer to the Tanium documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#eid_last_seen ZeroTrustDevicePostureRule#eid_last_seen}
	EidLastSeen *string `field:"optional" json:"eidLastSeen" yaml:"eidLastSeen"`
	// Enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#enabled ZeroTrustDevicePostureRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether or not file exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#exists ZeroTrustDevicePostureRule#exists}
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// List of values indicating purposes for which the certificate public key can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#extended_key_usage ZeroTrustDevicePostureRule#extended_key_usage}
	ExtendedKeyUsage *[]*string `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// List ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#id ZeroTrustDevicePostureRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether device is infected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#infected ZeroTrustDevicePostureRule#infected}
	Infected interface{} `field:"optional" json:"infected" yaml:"infected"`
	// Whether device is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#is_active ZeroTrustDevicePostureRule#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The Number of Issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#issue_count ZeroTrustDevicePostureRule#issue_count}
	IssueCount *string `field:"optional" json:"issueCount" yaml:"issueCount"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#last_seen ZeroTrustDevicePostureRule#last_seen}
	LastSeen *string `field:"optional" json:"lastSeen" yaml:"lastSeen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#locations ZeroTrustDevicePostureRule#locations}.
	Locations *ZeroTrustDevicePostureRuleInputLocations `field:"optional" json:"locations" yaml:"locations"`
	// Network status of device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#network_status ZeroTrustDevicePostureRule#network_status}
	NetworkStatus *string `field:"optional" json:"networkStatus" yaml:"networkStatus"`
	// Operating system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#operating_system ZeroTrustDevicePostureRule#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// Agent operational state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#operational_state ZeroTrustDevicePostureRule#operational_state}
	OperationalState *string `field:"optional" json:"operationalState" yaml:"operationalState"`
	// operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#operator ZeroTrustDevicePostureRule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Os Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#os ZeroTrustDevicePostureRule#os}
	Os *string `field:"optional" json:"os" yaml:"os"`
	// Operating System Distribution Name (linux only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#os_distro_name ZeroTrustDevicePostureRule#os_distro_name}
	OsDistroName *string `field:"optional" json:"osDistroName" yaml:"osDistroName"`
	// Version of OS Distribution (linux only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#os_distro_revision ZeroTrustDevicePostureRule#os_distro_revision}
	OsDistroRevision *string `field:"optional" json:"osDistroRevision" yaml:"osDistroRevision"`
	// Additional version data.
	//
	// For Mac or iOS, the Product Version Extra. For Linux, the kernel release version. (Mac, iOS, and Linux only)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#os_version_extra ZeroTrustDevicePostureRule#os_version_extra}
	OsVersionExtra *string `field:"optional" json:"osVersionExtra" yaml:"osVersionExtra"`
	// overall.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#overall ZeroTrustDevicePostureRule#overall}
	Overall *string `field:"optional" json:"overall" yaml:"overall"`
	// File path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#path ZeroTrustDevicePostureRule#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether to check all disks for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#require_all ZeroTrustDevicePostureRule#require_all}
	RequireAll interface{} `field:"optional" json:"requireAll" yaml:"requireAll"`
	// For more details on risk level, refer to the Tanium documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#risk_level ZeroTrustDevicePostureRule#risk_level}
	RiskLevel *string `field:"optional" json:"riskLevel" yaml:"riskLevel"`
	// A value between 0-100 assigned to devices set by the 3rd party posture provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#score ZeroTrustDevicePostureRule#score}
	Score *float64 `field:"optional" json:"score" yaml:"score"`
	// Score Operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#score_operator ZeroTrustDevicePostureRule#score_operator}
	ScoreOperator *string `field:"optional" json:"scoreOperator" yaml:"scoreOperator"`
	// SensorConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#sensor_config ZeroTrustDevicePostureRule#sensor_config}
	SensorConfig *string `field:"optional" json:"sensorConfig" yaml:"sensorConfig"`
	// SHA-256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#sha256 ZeroTrustDevicePostureRule#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
	// For more details on state, please refer to the Crowdstrike documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#state ZeroTrustDevicePostureRule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Signing certificate thumbprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#thumbprint ZeroTrustDevicePostureRule#thumbprint}
	Thumbprint *string `field:"optional" json:"thumbprint" yaml:"thumbprint"`
	// For more details on total score, refer to the Tanium documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#total_score ZeroTrustDevicePostureRule#total_score}
	TotalScore *float64 `field:"optional" json:"totalScore" yaml:"totalScore"`
	// Version of OS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#version ZeroTrustDevicePostureRule#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Version Operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.1.0/docs/resources/zero_trust_device_posture_rule#version_operator ZeroTrustDevicePostureRule#version_operator}
	VersionOperator *string `field:"optional" json:"versionOperator" yaml:"versionOperator"`
}

