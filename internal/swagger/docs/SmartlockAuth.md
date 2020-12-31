# SmartlockAuth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id for the smartlock authorization | [default to null]
**SmartlockId** | **int64** | The smartlock id | [default to null]
**AccountUserId** | **int32** | The id of the linked account user | [optional] [default to null]
**AuthId** | **int32** | The smartlock authorization id | [optional] [default to null]
**Code** | **int32** | The keypad code (only for type keypad) | [optional] [default to null]
**Type_** | **int32** | The type of the authorization: 0 .. app, 1 .. bridge, 2 .. fob, 3 .. keypad, 13 .. keypad code, 14 .. z-key, 15 .. virtual | [default to null]
**Name** | **string** | The name of the authorization (max 32 chars) | [default to null]
**Enabled** | **bool** | True if the auth is enabled | [default to null]
**RemoteAllowed** | **bool** | True if the auth has remote access | [default to null]
**LockCount** | **int32** | The lock count | [default to null]
**AllowedFromDate** | [**time.Time**](time.Time.md) | The allowed from date | [optional] [default to null]
**AllowedUntilDate** | [**time.Time**](time.Time.md) | The allowed until date | [optional] [default to null]
**AllowedWeekDays** | **int32** | The allowed weekdays bitmask: 64 .. monday, 32 .. tuesday, 16 .. wednesday, 8 .. thursday, 4 .. friday, 2 .. saturday, 1 .. sunday | [optional] [default to null]
**AllowedFromTime** | **int32** | The allowed from time (in minutes from midnight) | [optional] [default to null]
**AllowedUntilTime** | **int32** | The allowed until time (in minutes from midnight) | [optional] [default to null]
**LastActiveDate** | [**time.Time**](time.Time.md) | The last active date | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The creation date | [optional] [default to null]
**UpdateDate** | [**time.Time**](time.Time.md) | The update date | [optional] [default to null]
**OperationId** | [***ObjectId**](ObjectId.md) | The operation id - if set the auth is locked for another operations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


