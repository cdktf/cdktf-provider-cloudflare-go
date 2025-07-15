// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waitingroom

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaitingRoomConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The host name to which the waiting room will be applied (no wildcards).
	//
	// Please do not include the scheme (http:// or https://). The host and path combination must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#host WaitingRoom#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// A unique name to identify the waiting room. Only alphanumeric characters, hyphens and underscores are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#name WaitingRoom#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the number of new users that will be let into the route every minute.
	//
	// This value is used as baseline for the number of users that are let in per minute. So it is possible that there is a little more or little less traffic coming to the route based on the traffic patterns at that time around the world.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#new_users_per_minute WaitingRoom#new_users_per_minute}
	NewUsersPerMinute *float64 `field:"required" json:"newUsersPerMinute" yaml:"newUsersPerMinute"`
	// Sets the total number of active user sessions on the route at a point in time.
	//
	// A route is a combination of host and path on which a waiting room is available. This value is used as a baseline for the total number of active user sessions on the route. It is possible to have a situation where there are more or less active users sessions on the route based on the traffic patterns at that time around the world.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#total_active_users WaitingRoom#total_active_users}
	TotalActiveUsers *float64 `field:"required" json:"totalActiveUsers" yaml:"totalActiveUsers"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#zone_id WaitingRoom#zone_id}
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// Only available for the Waiting Room Advanced subscription.
	//
	// Additional hostname and path combinations to which this waiting room will be applied. There is an implied wildcard at the end of the path. The hostname and path combination must be unique to this and all other waiting rooms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#additional_routes WaitingRoom#additional_routes}
	AdditionalRoutes interface{} `field:"optional" json:"additionalRoutes" yaml:"additionalRoutes"`
	// Configures cookie attributes for the waiting room cookie.
	//
	// This encrypted cookie stores a user's status in the waiting room, such as queue position.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#cookie_attributes WaitingRoom#cookie_attributes}
	CookieAttributes *WaitingRoomCookieAttributes `field:"optional" json:"cookieAttributes" yaml:"cookieAttributes"`
	// Appends a '_' + a custom suffix to the end of Cloudflare Waiting Room's cookie name(__cf_waitingroom).
	//
	// If `cookie_suffix` is "abcd", the cookie name will be `__cf_waitingroom_abcd`. This field is required if using `additional_routes`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#cookie_suffix WaitingRoom#cookie_suffix}
	CookieSuffix *string `field:"optional" json:"cookieSuffix" yaml:"cookieSuffix"`
	// Only available for the Waiting Room Advanced subscription.
	//
	// This is a template html file that will be rendered at the edge. If no custom_page_html is provided, the default waiting room will be used. The template is based on mustache ( https://mustache.github.io/ ). There are several variables that are evaluated by the Cloudflare edge:
	// 1. {{`waitTimeKnown`}} Acts like a boolean value that indicates the behavior to take when wait time is not available, for instance when queue_all is **true**.
	// 2. {{`waitTimeFormatted`}} Estimated wait time for the user. For example, five minutes. Alternatively, you can use:
	// 3. {{`waitTime`}} Number of minutes of estimated wait for a user.
	// 4. {{`waitTimeHours`}} Number of hours of estimated wait for a user (`Math.floor(waitTime/60)`).
	// 5. {{`waitTimeHourMinutes`}} Number of minutes above the `waitTimeHours` value (`waitTime%60`).
	// 6. {{`queueIsFull`}} Changes to **true** when no more people can be added to the queue.
	//
	// To view the full list of variables, look at the `cfWaitingRoom` object described under the `json_response_enabled` property in other Waiting Room API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#custom_page_html WaitingRoom#custom_page_html}
	CustomPageHtml *string `field:"optional" json:"customPageHtml" yaml:"customPageHtml"`
	// The language of the default page template.
	//
	// If no default_template_language is provided, then `en-US` (English) will be used.
	// Available values: "en-US", "es-ES", "de-DE", "fr-FR", "it-IT", "ja-JP", "ko-KR", "pt-BR", "zh-CN", "zh-TW", "nl-NL", "pl-PL", "id-ID", "tr-TR", "ar-EG", "ru-RU", "fa-IR", "bg-BG", "hr-HR", "cs-CZ", "da-DK", "fi-FI", "lt-LT", "ms-MY", "nb-NO", "ro-RO", "el-GR", "he-IL", "hi-IN", "hu-HU", "sr-BA", "sk-SK", "sl-SI", "sv-SE", "tl-PH", "th-TH", "uk-UA", "vi-VN".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#default_template_language WaitingRoom#default_template_language}
	DefaultTemplateLanguage *string `field:"optional" json:"defaultTemplateLanguage" yaml:"defaultTemplateLanguage"`
	// A note that you can use to add more details about the waiting room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#description WaitingRoom#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Only available for the Waiting Room Advanced subscription.
	//
	// Disables automatic renewal of session cookies. If `true`, an accepted user will have session_duration minutes to browse the site. After that, they will have to go through the waiting room again. If `false`, a user's session cookie will be automatically renewed on every request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#disable_session_renewal WaitingRoom#disable_session_renewal}
	DisableSessionRenewal interface{} `field:"optional" json:"disableSessionRenewal" yaml:"disableSessionRenewal"`
	// A list of enabled origin commands.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#enabled_origin_commands WaitingRoom#enabled_origin_commands}
	EnabledOriginCommands *[]*string `field:"optional" json:"enabledOriginCommands" yaml:"enabledOriginCommands"`
	// Only available for the Waiting Room Advanced subscription.
	//
	// If `true`, requests to the waiting room with the header `Accept: application/json` will receive a JSON response object with information on the user's status in the waiting room as opposed to the configured static HTML page. This JSON response object has one property `cfWaitingRoom` which is an object containing the following fields:
	// 1. `inWaitingRoom`: Boolean indicating if the user is in the waiting room (always **true**).
	// 2. `waitTimeKnown`: Boolean indicating if the current estimated wait times are accurate. If **false**, they are not available.
	// 3. `waitTime`: Valid only when `waitTimeKnown` is **true**. Integer indicating the current estimated time in minutes the user will wait in the waiting room. When `queueingMethod` is **random**, this is set to `waitTime50Percentile`.
	// 4. `waitTime25Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 25% of users that gain entry the fastest (25th percentile).
	// 5. `waitTime50Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 50% of users that gain entry the fastest (50th percentile). In other words, half of the queued users are expected to let into the origin website before `waitTime50Percentile` and half are expected to be let in after it.
	// 6. `waitTime75Percentile`: Valid only when `queueingMethod` is **random** and `waitTimeKnown` is **true**. Integer indicating the current estimated maximum wait time for the 75% of users that gain entry the fastest (75th percentile).
	// 7. `waitTimeFormatted`: String displaying the `waitTime` formatted in English for users. If `waitTimeKnown` is **false**, `waitTimeFormatted` will display **unavailable**.
	// 8. `queueIsFull`: Boolean indicating if the waiting room's queue is currently full and not accepting new users at the moment.
	// 9. `queueAll`: Boolean indicating if all users will be queued in the waiting room and no one will be let into the origin website.
	// 10. `lastUpdated`: String displaying the timestamp as an ISO 8601 string of the user's last attempt to leave the waiting room and be let into the origin website. The user is able to make another attempt after `refreshIntervalSeconds` past this time. If the user makes a request too soon, it will be ignored and `lastUpdated` will not change.
	// 11. `refreshIntervalSeconds`: Integer indicating the number of seconds after `lastUpdated` until the user is able to make another attempt to leave the waiting room and be let into the origin website. When the `queueingMethod` is `reject`, there is no specified refresh time —\_it will always be **zero**.
	// 12. `queueingMethod`: The queueing method currently used by the waiting room. It is either **fifo**, **random**, **passthrough**, or **reject**.
	// 13. `isFIFOQueue`: Boolean indicating if the waiting room uses a FIFO (First-In-First-Out) queue.
	// 14. `isRandomQueue`: Boolean indicating if the waiting room uses a Random queue where users gain access randomly.
	// 15. `isPassthroughQueue`: Boolean indicating if the waiting room uses a passthrough queue. Keep in mind that when passthrough is enabled, this JSON response will only exist when `queueAll` is **true** or `isEventPrequeueing` is **true** because in all other cases requests will go directly to the origin.
	// 16. `isRejectQueue`: Boolean indicating if the waiting room uses a reject queue.
	// 17. `isEventActive`: Boolean indicating if an event is currently occurring. Events are able to change a waiting room's behavior during a specified period of time. For additional information, look at the event properties `prequeue_start_time`, `event_start_time`, and `event_end_time` in the documentation for creating waiting room events. Events are considered active between these start and end times, as well as during the prequeueing period if it exists.
	// 18. `isEventPrequeueing`: Valid only when `isEventActive` is **true**. Boolean indicating if an event is currently prequeueing users before it starts.
	// 19. `timeUntilEventStart`: Valid only when `isEventPrequeueing` is **true**. Integer indicating the number of minutes until the event starts.
	// 20. `timeUntilEventStartFormatted`: String displaying the `timeUntilEventStart` formatted in English for users. If `isEventPrequeueing` is **false**, `timeUntilEventStartFormatted` will display **unavailable**.
	// 21. `timeUntilEventEnd`: Valid only when `isEventActive` is **true**. Integer indicating the number of minutes until the event ends.
	// 22. `timeUntilEventEndFormatted`: String displaying the `timeUntilEventEnd` formatted in English for users. If `isEventActive` is **false**, `timeUntilEventEndFormatted` will display **unavailable**.
	// 23. `shuffleAtEventStart`: Valid only when `isEventActive` is **true**. Boolean indicating if the users in the prequeue are shuffled randomly when the event starts.
	// 24. `turnstile`: Empty when turnstile isn't enabled. String displaying an html tag to display the Turnstile widget. Please add the `{{{turnstile}}}` tag to the `custom_html` template to ensure the Turnstile widget appears.
	// 25. `infiniteQueue`: Boolean indicating whether the response is for a user in the infinite queue.
	//
	// An example cURL to a waiting room could be:
	//
	// 	curl -X GET "https://example.com/waitingroom" \
	// 		-H "Accept: application/json"
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an example JSON response when `queueingMethod` is **fifo** and no event is active could be:
	//
	// 	{
	// 		"cfWaitingRoom": {
	// 			"inWaitingRoom": true,
	// 			"waitTimeKnown": true,
	// 			"waitTime": 10,
	// 			"waitTime25Percentile": 0,
	// 			"waitTime50Percentile": 0,
	// 			"waitTime75Percentile": 0,
	// 			"waitTimeFormatted": "10 minutes",
	// 			"queueIsFull": false,
	// 			"queueAll": false,
	// 			"lastUpdated": "2020-08-03T23:46:00.000Z",
	// 			"refreshIntervalSeconds": 20,
	// 			"queueingMethod": "fifo",
	// 			"isFIFOQueue": true,
	// 			"isRandomQueue": false,
	// 			"isPassthroughQueue": false,
	// 			"isRejectQueue": false,
	// 			"isEventActive": false,
	// 			"isEventPrequeueing": false,
	// 			"timeUntilEventStart": 0,
	// 			"timeUntilEventStartFormatted": "unavailable",
	// 			"timeUntilEventEnd": 0,
	// 			"timeUntilEventEndFormatted": "unavailable",
	// 			"shuffleAtEventStart": false
	// 		}
	// 	}
	//
	// If `json_response_enabled` is **true** and the request hits the waiting room, an example JSON response when `queueingMethod` is **random** and an event is active could be:
	//
	// 	{
	// 		"cfWaitingRoom": {
	// 			"inWaitingRoom": true,
	// 			"waitTimeKnown": true,
	// 			"waitTime": 10,
	// 			"waitTime25Percentile": 5,
	// 			"waitTime50Percentile": 10,
	// 			"waitTime75Percentile": 15,
	// 			"waitTimeFormatted": "5 minutes to 15 minutes",
	// 			"queueIsFull": false,
	// 			"queueAll": false,
	// 			"lastUpdated": "2020-08-03T23:46:00.000Z",
	// 			"refreshIntervalSeconds": 20,
	// 			"queueingMethod": "random",
	// 			"isFIFOQueue": false,
	// 			"isRandomQueue": true,
	// 			"isPassthroughQueue": false,
	// 			"isRejectQueue": false,
	// 			"isEventActive": true,
	// 			"isEventPrequeueing": false,
	// 			"timeUntilEventStart": 0,
	// 			"timeUntilEventStartFormatted": "unavailable",
	// 			"timeUntilEventEnd": 15,
	// 			"timeUntilEventEndFormatted": "15 minutes",
	// 			"shuffleAtEventStart": true
	// 		}
	// 	}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#json_response_enabled WaitingRoom#json_response_enabled}
	JsonResponseEnabled interface{} `field:"optional" json:"jsonResponseEnabled" yaml:"jsonResponseEnabled"`
	// Sets the path within the host to enable the waiting room on.
	//
	// The waiting room will be enabled for all subpaths as well. If there are two waiting rooms on the same subpath, the waiting room for the most specific path will be chosen. Wildcards and query parameters are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#path WaitingRoom#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If queue_all is `true`, all the traffic that is coming to a route will be sent to the waiting room.
	//
	// No new traffic can get to the route once this field is set and estimated time will become unavailable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#queue_all WaitingRoom#queue_all}
	QueueAll interface{} `field:"optional" json:"queueAll" yaml:"queueAll"`
	// Sets the queueing method used by the waiting room.
	//
	// Changing this parameter from the **default** queueing method is only available for the Waiting Room Advanced subscription. Regardless of the queueing method, if `queue_all` is enabled or an event is prequeueing, users in the waiting room will not be accepted to the origin. These users will always see a waiting room page that refreshes automatically. The valid queueing methods are:
	// 1. `fifo` **(default)**: First-In-First-Out queue where customers gain access in the order they arrived.
	// 2. `random`: Random queue where customers gain access randomly, regardless of arrival time.
	// 3. `passthrough`: Users will pass directly through the waiting room and into the origin website. As a result, any configured limits will not be respected while this is enabled. This method can be used as an alternative to disabling a waiting room (with `suspended`) so that analytics are still reported. This can be used if you wish to allow all traffic normally, but want to restrict traffic during a waiting room event, or vice versa.
	// 4. `reject`: Users will be immediately rejected from the waiting room. As a result, no users will reach the origin website while this is enabled. This can be used if you wish to reject all traffic while performing maintenance, block traffic during a specified period of time (an event), or block traffic while events are not occurring. Consider a waiting room used for vaccine distribution that only allows traffic during sign-up events, and otherwise blocks all traffic. For this case, the waiting room uses `reject`, and its events override this with `fifo`, `random`, or `passthrough`. When this queueing method is enabled and neither `queueAll` is enabled nor an event is prequeueing, the waiting room page **will not refresh automatically**.
	// Available values: "fifo", "random", "passthrough", "reject".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#queueing_method WaitingRoom#queueing_method}
	QueueingMethod *string `field:"optional" json:"queueingMethod" yaml:"queueingMethod"`
	// HTTP status code returned to a user while in the queue. Available values: 200, 202, 429.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#queueing_status_code WaitingRoom#queueing_status_code}
	QueueingStatusCode *float64 `field:"optional" json:"queueingStatusCode" yaml:"queueingStatusCode"`
	// Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the route.
	//
	// If a user is not seen by Cloudflare again in that time period, they will be treated as a new user that visits the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#session_duration WaitingRoom#session_duration}
	SessionDuration *float64 `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// Suspends or allows traffic going to the waiting room.
	//
	// If set to `true`, the traffic will not go to the waiting room.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#suspended WaitingRoom#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// Which action to take when a bot is detected using Turnstile.
	//
	// `log` will
	// have no impact on queueing behavior, simply keeping track of how many
	// bots are detected in Waiting Room Analytics. `infinite_queue` will send
	// bots to a false queueing state, where they will never reach your
	// origin. `infinite_queue` requires Advanced Waiting Room.
	// Available values: "log", "infinite_queue".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#turnstile_action WaitingRoom#turnstile_action}
	TurnstileAction *string `field:"optional" json:"turnstileAction" yaml:"turnstileAction"`
	// Which Turnstile widget type to use for detecting bot traffic.
	//
	// See
	// [the Turnstile documentation](https://developers.cloudflare.com/turnstile/concepts/widget/#widget-types)
	// for the definitions of these widget types. Set to `off` to disable the
	// Turnstile integration entirely. Setting this to anything other than
	// `off` or `invisible` requires Advanced Waiting Room.
	// Available values: "off", "invisible", "visible_non_interactive", "visible_managed".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.0/docs/resources/waiting_room#turnstile_mode WaitingRoom#turnstile_mode}
	TurnstileMode *string `field:"optional" json:"turnstileMode" yaml:"turnstileMode"`
}

