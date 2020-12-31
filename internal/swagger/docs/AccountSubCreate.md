# AccountSubCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address | [default to null]
**Password** | **string** | The password (must be at least 5 chars long) | [default to null]
**Name** | **string** | The name of the sub account | [default to null]
**Rights** | **int32** | The right bitmask of the sub account: 1 .. operate smartlock, 2 .. change smartlock config, 4 .. manage smartlock users, 8 .. view smartlock logs, 16 .. manage sub accounts | [default to null]
**Language** | **string** | The language code | [default to null]
**Profile** | [***AccountProfile**](Account.Profile.md) | The optional profile | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


