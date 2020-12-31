# SmartlockLog

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id for the smartlock log | [default to null]
**SmartlockId** | **int64** | The smartlock id | [default to null]
**DeviceType** | **int32** | The device type: 0 .. smartlock and box, 2 .. opener, 3 .. smartdoor | [default to null]
**AccountUserId** | **int32** | The id of the linked account user | [optional] [default to null]
**AuthId** | **string** | The id of the linked smartlock auth | [optional] [default to null]
**Name** | **string** | The name | [default to null]
**Action** | **int32** | The action: 1 .. unlock, 2 .. lock, 3 .. unlatch, 4 .. lock&#39;n&#39;go, 5 .. lock&#39;n&#39;go with unlatch, 208 .. door warning ajar, 209 door warning status mismatch, 224 .. doorbell recognition (only Opener), 240 .. door opened, 241 .. door closed, 242 .. door sensor jammed, 243 .. firmware update, 250 .. door log enabled, 251 .. door log disabled, 252 .. initialization, 253 .. calibration, 254 .. log enabled, 255 .. log disabled | [default to null]
**Trigger** | **int32** | The trigger: 0 .. system, 1 .. manual, 2 .. button, 3 .. automatic, 4 .. web, 5 .. app, 6 .. auto lock, 7 .. accessory, 255 .. keypad | [default to null]
**State** | **int32** | The completion state: 0 .. Success, 1 .. Motor blocked, 2 .. Canceled, 3 .. Too recent, 4 .. Busy, 5 .. Low motor voltage, 6 .. Clutch failure, 7 .. Motor power failure, 8 .. Incomplete, 9 .. Rejected, 10 .. Rejected night mode, 254 .. Other error, 255 .. Unknown error | [default to null]
**AutoUnlock** | **bool** | True if it was an auto unlock | [default to null]
**Date** | [**time.Time**](time.Time.md) | The log date | [default to null]
**OpenerLog** | [***SmartlockLogOpenerLog**](SmartlockLog.OpenerLog.md) | The opener specific log | [optional] [default to null]
**AjarTimeout** | **int32** | The door sensor warning ajar timeout (in minutes, only for action &#x3D; 208) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


