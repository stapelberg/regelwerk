# SmartlocksAuthAdvancedCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the authorization (max 32 chars) | [default to null]
**AllowedFromDate** | [**time.Time**](time.Time.md) | The allowed from date | [optional] [default to null]
**AllowedUntilDate** | [**time.Time**](time.Time.md) | The allowed until date | [optional] [default to null]
**AllowedWeekDays** | **int32** | The allowed weekdays bitmask: 64 .. monday, 32 .. tuesday, 16 .. wednesday, 8 .. thursday, 4 .. friday, 2 .. saturday, 1 .. sunday | [optional] [default to null]
**AllowedFromTime** | **int32** | The allowed from time (in minutes from midnight) | [optional] [default to null]
**AllowedUntilTime** | **int32** | The allowed until time (in minutes from midnight) | [optional] [default to null]
**AccountUserId** | **int32** | The id of the linked account user | [default to null]
**SmartlockIds** | **[]int64** | The list of smartlock ids | [default to null]
**RemoteAllowed** | **bool** | True if the auth has remote access | [default to null]
**SmartActionsEnabled** | **bool** | The smart actions enabled flag | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


