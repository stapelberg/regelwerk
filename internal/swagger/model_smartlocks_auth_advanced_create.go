/*
 * Nuki API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Contact: contact@nuki.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type SmartlocksAuthAdvancedCreate struct {
	// The name of the authorization (max 32 chars)
	Name string `json:"name"`
	// The allowed from date
	AllowedFromDate time.Time `json:"allowedFromDate,omitempty"`
	// The allowed until date
	AllowedUntilDate time.Time `json:"allowedUntilDate,omitempty"`
	// The allowed weekdays bitmask: 64 .. monday, 32 .. tuesday, 16 .. wednesday, 8 .. thursday, 4 .. friday, 2 .. saturday, 1 .. sunday
	AllowedWeekDays int32 `json:"allowedWeekDays,omitempty"`
	// The allowed from time (in minutes from midnight)
	AllowedFromTime int32 `json:"allowedFromTime,omitempty"`
	// The allowed until time (in minutes from midnight)
	AllowedUntilTime int32 `json:"allowedUntilTime,omitempty"`
	// The id of the linked account user
	AccountUserId int32 `json:"accountUserId"`
	// The list of smartlock ids
	SmartlockIds []int64 `json:"smartlockIds"`
	// True if the auth has remote access
	RemoteAllowed bool `json:"remoteAllowed"`
	// The smart actions enabled flag
	SmartActionsEnabled bool `json:"smartActionsEnabled,omitempty"`
}
