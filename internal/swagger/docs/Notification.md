# Notification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | **string** | The unique notificationId for the notification | [default to null]
**ReferenceId** | **string** | The reference ID, an ID to identify a foreign system | [optional] [default to null]
**PushId** | **string** | The push ID or the POST URL for a webhook | [default to null]
**Secret** | **string** | The 40 byte hex string to sign the checksumof the POST payload if the notification is webhook (os&#x3D;2) | [optional] [default to null]
**Os** | **int32** | The operating system: 0 .. Android, 1 .. iOS, 2 .. web hook | [default to null]
**Language** | **string** | The language of push messages: cs, de, en (default), es, fr, it, nl, sk | [optional] [default to null]
**Status** | **int32** | Current state: 0 .. init, 1 .. active, 2 .. failed | [optional] [default to null]
**LastActiveDate** | [**time.Time**](time.Time.md) | The last active date | [optional] [default to null]
**Settings** | [**[]NotificationSetting**](Notification.Setting.md) | Settings per Smart Lock | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


