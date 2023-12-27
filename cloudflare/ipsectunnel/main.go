// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ipsectunnel

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		reflect.TypeOf((*IpsecTunnel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowNullCipher", GoGetter: "AllowNullCipher"},
			_jsii_.MemberProperty{JsiiProperty: "allowNullCipherInput", GoGetter: "AllowNullCipherInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudflareEndpoint", GoGetter: "CloudflareEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "cloudflareEndpointInput", GoGetter: "CloudflareEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerEndpoint", GoGetter: "CustomerEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "customerEndpointInput", GoGetter: "CustomerEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqdnId", GoGetter: "FqdnId"},
			_jsii_.MemberProperty{JsiiProperty: "fqdnIdInput", GoGetter: "FqdnIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckEnabled", GoGetter: "HealthCheckEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckEnabledInput", GoGetter: "HealthCheckEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTarget", GoGetter: "HealthCheckTarget"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTargetInput", GoGetter: "HealthCheckTargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckType", GoGetter: "HealthCheckType"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTypeInput", GoGetter: "HealthCheckTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "hexId", GoGetter: "HexId"},
			_jsii_.MemberProperty{JsiiProperty: "hexIdInput", GoGetter: "HexIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceAddress", GoGetter: "InterfaceAddress"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceAddressInput", GoGetter: "InterfaceAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "psk", GoGetter: "Psk"},
			_jsii_.MemberProperty{JsiiProperty: "pskInput", GoGetter: "PskInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteId", GoGetter: "RemoteId"},
			_jsii_.MemberProperty{JsiiProperty: "remoteIdInput", GoGetter: "RemoteIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowNullCipher", GoMethod: "ResetAllowNullCipher"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetFqdnId", GoMethod: "ResetFqdnId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckEnabled", GoMethod: "ResetHealthCheckEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckTarget", GoMethod: "ResetHealthCheckTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckType", GoMethod: "ResetHealthCheckType"},
			_jsii_.MemberMethod{JsiiMethod: "resetHexId", GoMethod: "ResetHexId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPsk", GoMethod: "ResetPsk"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteId", GoMethod: "ResetRemoteId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserId", GoMethod: "ResetUserId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userIdInput", GoGetter: "UserIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_IpsecTunnel{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnelConfig",
		reflect.TypeOf((*IpsecTunnelConfig)(nil)).Elem(),
	)
}
