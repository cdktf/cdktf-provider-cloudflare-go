// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package zerotrustdeviceposturerule


type ZeroTrustDevicePostureRuleInput struct {
	// The number of active threats from SentinelOne.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#active_threats ZeroTrustDevicePostureRule#active_threats}
	ActiveThreats *float64 `field:"optional" json:"activeThreats" yaml:"activeThreats"`
	// The UUID of a Cloudflare managed certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#certificate_id ZeroTrustDevicePostureRule#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// Specific volume(s) to check for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#check_disks ZeroTrustDevicePostureRule#check_disks}
	CheckDisks *[]*string `field:"optional" json:"checkDisks" yaml:"checkDisks"`
	// Confirm the certificate was not imported from another device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#check_private_key ZeroTrustDevicePostureRule#check_private_key}
	CheckPrivateKey interface{} `field:"optional" json:"checkPrivateKey" yaml:"checkPrivateKey"`
	// The common name for a certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#cn ZeroTrustDevicePostureRule#cn}
	Cn *string `field:"optional" json:"cn" yaml:"cn"`
	// The workspace one or intune device compliance status.
	//
	// `compliant` and `noncompliant` are values supported by both providers. `unknown`, `conflict`, `error`, `ingraceperiod` values are only supported by intune. Available values: `compliant`, `noncompliant`, `unknown`, `conflict`, `error`, `ingraceperiod`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#compliance_status ZeroTrustDevicePostureRule#compliance_status}
	ComplianceStatus *string `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// The workspace one or intune connection id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#connection_id ZeroTrustDevicePostureRule#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The count comparison operator for kolide. Available values: `>`, `>=`, `<`, `<=`, `==`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#count_operator ZeroTrustDevicePostureRule#count_operator}
	CountOperator *string `field:"optional" json:"countOperator" yaml:"countOperator"`
	// The domain that the client must join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#domain ZeroTrustDevicePostureRule#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The time a device last seen in Tanium.
	//
	// Must be in the format `1h` or `30m`. Valid units are `d`, `h` and `m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#eid_last_seen ZeroTrustDevicePostureRule#eid_last_seen}
	EidLastSeen *string `field:"optional" json:"eidLastSeen" yaml:"eidLastSeen"`
	// True if the firewall must be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#enabled ZeroTrustDevicePostureRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Checks if the file should exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#exists ZeroTrustDevicePostureRule#exists}
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// List of values indicating purposes for which the certificate public key can be used. Available values: `clientAuth`, `emailProtection`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#extended_key_usage ZeroTrustDevicePostureRule#extended_key_usage}
	ExtendedKeyUsage *[]*string `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// The Teams List id. Required for `serial_number` and `unique_client_id` rule types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#id ZeroTrustDevicePostureRule#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// True if SentinelOne device is infected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#infected ZeroTrustDevicePostureRule#infected}
	Infected interface{} `field:"optional" json:"infected" yaml:"infected"`
	// True if SentinelOne device is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#is_active ZeroTrustDevicePostureRule#is_active}
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The number of issues for kolide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#issue_count ZeroTrustDevicePostureRule#issue_count}
	IssueCount *string `field:"optional" json:"issueCount" yaml:"issueCount"`
	// The duration of time that the host was last seen from Crowdstrike.
	//
	// Must be in the format `1h` or `30m`. Valid units are `d`, `h` and `m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#last_seen ZeroTrustDevicePostureRule#last_seen}
	LastSeen *string `field:"optional" json:"lastSeen" yaml:"lastSeen"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#locations ZeroTrustDevicePostureRule#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// The network status from SentinelOne. Available values: `connected`, `disconnected`, `disconnecting`, `connecting`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#network_status ZeroTrustDevicePostureRule#network_status}
	NetworkStatus *string `field:"optional" json:"networkStatus" yaml:"networkStatus"`
	// The version comparison operator. Available values: `>`, `>=`, `<`, `<=`, `==`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#operator ZeroTrustDevicePostureRule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// OS signal score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#os ZeroTrustDevicePostureRule#os}
	Os *string `field:"optional" json:"os" yaml:"os"`
	// The operating system excluding version information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#os_distro_name ZeroTrustDevicePostureRule#os_distro_name}
	OsDistroName *string `field:"optional" json:"osDistroName" yaml:"osDistroName"`
	// The operating system version excluding OS name information or release name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#os_distro_revision ZeroTrustDevicePostureRule#os_distro_revision}
	OsDistroRevision *string `field:"optional" json:"osDistroRevision" yaml:"osDistroRevision"`
	// Extra version value following the operating system semantic version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#os_version_extra ZeroTrustDevicePostureRule#os_version_extra}
	OsVersionExtra *string `field:"optional" json:"osVersionExtra" yaml:"osVersionExtra"`
	// Overall ZTA score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#overall ZeroTrustDevicePostureRule#overall}
	Overall *string `field:"optional" json:"overall" yaml:"overall"`
	// The path to the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#path ZeroTrustDevicePostureRule#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// True if all drives must be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#require_all ZeroTrustDevicePostureRule#require_all}
	RequireAll interface{} `field:"optional" json:"requireAll" yaml:"requireAll"`
	// The risk level from Tanium. Available values: `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#risk_level ZeroTrustDevicePostureRule#risk_level}
	RiskLevel *string `field:"optional" json:"riskLevel" yaml:"riskLevel"`
	// Checks if the application should be running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#running ZeroTrustDevicePostureRule#running}
	Running interface{} `field:"optional" json:"running" yaml:"running"`
	// Sensor signal score from Crowdstrike. Value must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#sensor_config ZeroTrustDevicePostureRule#sensor_config}
	SensorConfig *string `field:"optional" json:"sensorConfig" yaml:"sensorConfig"`
	// The sha256 hash of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#sha256 ZeroTrustDevicePostureRule#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
	// The host’s current online status from Crowdstrike. Available values: `online`, `offline`, `unknown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#state ZeroTrustDevicePostureRule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// The thumbprint of the file certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#thumbprint ZeroTrustDevicePostureRule#thumbprint}
	Thumbprint *string `field:"optional" json:"thumbprint" yaml:"thumbprint"`
	// The total score from Tanium.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#total_score ZeroTrustDevicePostureRule#total_score}
	TotalScore *float64 `field:"optional" json:"totalScore" yaml:"totalScore"`
	// The operating system semantic version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#version ZeroTrustDevicePostureRule#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The version comparison operator for crowdstrike. Available values: `>`, `>=`, `<`, `<=`, `==`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.41.0/docs/resources/zero_trust_device_posture_rule#version_operator ZeroTrustDevicePostureRule#version_operator}
	VersionOperator *string `field:"optional" json:"versionOperator" yaml:"versionOperator"`
}

