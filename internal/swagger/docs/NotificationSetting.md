# NotificationSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmartlockId** | **int64** | The smartlock ID, if not set all Smart Locks of the account  are enabled for push notifications | [optional] [default to null]
**TriggerEvents** | **[]string** | A set on which push notifications should be triggered: lock, unlock, unlatch, lockngo, open, ring, doorsensor, warnings, smartlock | [default to null]
**AuthIds** | **[]string** | A set of auth IDs to filter push notifications to certain  users or keypads. If no entry push notifications are triggered for all users and keypads | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


