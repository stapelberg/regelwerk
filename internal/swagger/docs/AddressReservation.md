# AddressReservation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id | [default to null]
**AddressId** | **int32** | The address id | [default to null]
**AccountId** | **int32** | The account id | [default to null]
**Email** | **string** | The email of the guest | [default to null]
**Name** | **string** | The name of the guest | [default to null]
**Guests** | **int32** | The number of guests | [default to null]
**GuestsIssued** | **int32** | The number of guests issued | [default to null]
**KeypadIssued** | **bool** | True if a keypad authorization was issued | [default to null]
**State** | **string** | The state | [default to null]
**ServiceId** | **string** | The optional service id if the address is from an partner service | [optional] [default to null]
**Reference** | **string** | The reference (booking code) | [optional] [default to null]
**Automation** | **int32** | The automation state | [default to null]
**CheckedIn** | **bool** | True if the user has checked in, false if the check in is pending, null if it isn&#39;t monitored | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The start date | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The end date | [default to null]
**UpdateDate** | [**time.Time**](time.Time.md) | The update date | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


