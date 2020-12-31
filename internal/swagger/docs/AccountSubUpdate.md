# AccountSubUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The new email address | [optional] [default to null]
**Password** | **string** | The new password (must be at least 5 chars long) | [optional] [default to null]
**Name** | **string** | The new name of the sub account | [optional] [default to null]
**Rights** | **int32** | The new right bitmask of the sub account: 1 .. operate smartlock, 2 .. change smartlock config, 4 .. manage smartlock users, 8 .. view smartlock logs, 16 .. manage sub accounts, 32 .. manage sub accounts, 64 .. create smartlocks | [optional] [default to null]
**Language** | **string** | The language code | [default to null]
**Config** | [***AccountConfig**](Account.Config.md) | The optional config | [optional] [default to null]
**Profile** | [***AccountProfile**](Account.Profile.md) | The optional profile | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


