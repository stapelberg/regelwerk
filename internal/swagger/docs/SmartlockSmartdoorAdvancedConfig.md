# SmartlockSmartdoorAdvancedConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LngTimeout** | **int32** | Timeout in seconds for lock ‘n’ go | [optional] [default to null]
**SingleButtonPressAction** | **int32** | The desired action, if the button is pressed once: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock &#39;n&#39; go, 6 .. show status | [optional] [default to null]
**DoubleButtonPressAction** | **int32** | The desired action, if the button is pressed twice: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock &#39;n&#39; go, 6 .. show status | [optional] [default to null]
**AutomaticBatteryTypeDetection** | **bool** | Flag that indicates if the automatic detection of the battery type is enabled | [optional] [default to null]
**UnlatchDuration** | **int32** | Duration in seconds for holding the latch in unlatched position | [optional] [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]
**BuzzerVolume** | **int32** | The volume of the buzzer: 0 .. off, 1 .. low, 2 .. normal | [optional] [default to null]
**SupportedBatteryTypes** | **[]int32** | Set of supported battery types: 0 .. alkali, 1 .. accumulator, 2 .. lithium, 3 .. fixed, 254 .. automatic, 255 .. unknown | [optional] [default to null]
**BatteryType** | **int32** | The type of the batteries present in the smart lock: 0 .. alkali, 1 .. accumulator, 2 .. lithium, 3 .. fixed, 255 .. unknown | [default to null]
**AutoLockTimeout** | **int32** | Seconds until the smart lock relocks itself after it has been unlocked. No auto relock if value is 0 | [optional] [default to null]
**AutoLock** | **bool** | The Auto Lock feature automatically locks your door when it has been unlocked for a certain period of time | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


