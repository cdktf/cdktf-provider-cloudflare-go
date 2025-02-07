// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroomevent

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.waitingRoomEvent.WaitingRoomEvent",
		reflect.TypeOf((*WaitingRoomEvent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdOn", GoGetter: "CreatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "customPageHtml", GoGetter: "CustomPageHtml"},
			_jsii_.MemberProperty{JsiiProperty: "customPageHtmlInput", GoGetter: "CustomPageHtmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableSessionRenewal", GoGetter: "DisableSessionRenewal"},
			_jsii_.MemberProperty{JsiiProperty: "disableSessionRenewalInput", GoGetter: "DisableSessionRenewalInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndTime", GoGetter: "EventEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "eventEndTimeInput", GoGetter: "EventEndTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventStartTime", GoGetter: "EventStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "eventStartTimeInput", GoGetter: "EventStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedOn", GoGetter: "ModifiedOn"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "newUsersPerMinute", GoGetter: "NewUsersPerMinute"},
			_jsii_.MemberProperty{JsiiProperty: "newUsersPerMinuteInput", GoGetter: "NewUsersPerMinuteInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "prequeueStartTime", GoGetter: "PrequeueStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "prequeueStartTimeInput", GoGetter: "PrequeueStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "queueingMethod", GoGetter: "QueueingMethod"},
			_jsii_.MemberProperty{JsiiProperty: "queueingMethodInput", GoGetter: "QueueingMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomPageHtml", GoMethod: "ResetCustomPageHtml"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableSessionRenewal", GoMethod: "ResetDisableSessionRenewal"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewUsersPerMinute", GoMethod: "ResetNewUsersPerMinute"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrequeueStartTime", GoMethod: "ResetPrequeueStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueueingMethod", GoMethod: "ResetQueueingMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionDuration", GoMethod: "ResetSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetShuffleAtEventStart", GoMethod: "ResetShuffleAtEventStart"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuspended", GoMethod: "ResetSuspended"},
			_jsii_.MemberMethod{JsiiMethod: "resetTotalActiveUsers", GoMethod: "ResetTotalActiveUsers"},
			_jsii_.MemberProperty{JsiiProperty: "sessionDuration", GoGetter: "SessionDuration"},
			_jsii_.MemberProperty{JsiiProperty: "sessionDurationInput", GoGetter: "SessionDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "shuffleAtEventStart", GoGetter: "ShuffleAtEventStart"},
			_jsii_.MemberProperty{JsiiProperty: "shuffleAtEventStartInput", GoGetter: "ShuffleAtEventStartInput"},
			_jsii_.MemberProperty{JsiiProperty: "suspended", GoGetter: "Suspended"},
			_jsii_.MemberProperty{JsiiProperty: "suspendedInput", GoGetter: "SuspendedInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalActiveUsers", GoGetter: "TotalActiveUsers"},
			_jsii_.MemberProperty{JsiiProperty: "totalActiveUsersInput", GoGetter: "TotalActiveUsersInput"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "waitingRoomId", GoGetter: "WaitingRoomId"},
			_jsii_.MemberProperty{JsiiProperty: "waitingRoomIdInput", GoGetter: "WaitingRoomIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneId", GoGetter: "ZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "zoneIdInput", GoGetter: "ZoneIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_WaitingRoomEvent{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.waitingRoomEvent.WaitingRoomEventConfig",
		reflect.TypeOf((*WaitingRoomEventConfig)(nil)).Elem(),
	)
}
