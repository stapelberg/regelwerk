# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int32** | The account id | [default to null]
**Type_** | **int32** | The type: 0 .. user, 1 .. company, 2 .. caretaker | [default to null]
**Email** | **string** | The email address | [default to null]
**Name** | **string** | The name | [default to null]
**MasterAccountId** | **int32** | The master account id if it&#39;s a sub account | [optional] [default to null]
**Rights** | **int32** | The rights bitmask if it&#39;s a sub account: 1 .. manage smartlock, 2 .. operate smartlock, 4 .. manage smartlock config, 8 .. manage smartlock authorizations, 16 .. view smartlock logs, 32 .. manage sub accounts, 64 .. create smartlocks | [optional] [default to null]
**Language** | **string** | The language code | [optional] [default to null]
**Config** | [***AccountConfig**](Account.Config.md) | The optional config | [optional] [default to null]
**Profile** | [***AccountProfile**](Account.Profile.md) | The optional profile | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The creation date | [default to null]
**UpdateDate** | [**time.Time**](time.Time.md) | The update date | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


