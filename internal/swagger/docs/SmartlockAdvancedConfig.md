# SmartlockAdvancedConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LngTimeout** | **int32** | Timeout in seconds for lock ‘n’ go | [optional] [default to null]
**SingleButtonPressAction** | **int32** | The desired action, if the button is pressed once: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock &#39;n&#39; go, 6 .. show status | [optional] [default to null]
**DoubleButtonPressAction** | **int32** | The desired action, if the button is pressed twice: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock &#39;n&#39; go, 6 .. show status | [optional] [default to null]
**AutomaticBatteryTypeDetection** | **bool** | Flag that indicates if the automatic detection of the battery type is enabled | [optional] [default to null]
**UnlatchDuration** | **int32** | Duration in seconds for holding the latch in unlatched position | [optional] [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]
**TotalDegrees** | **int32** | The absolute total position in degrees that has been reached during calibration | [default to null]
**SingleLockedPositionOffsetDegrees** | **int32** | Offset that alters the single locked position | [default to null]
**UnlockedToLockedTransitionOffsetDegrees** | **int32** | Offset that alters the position where transition from unlocked to locked happens | [optional] [default to null]
**UnlockedPositionOffsetDegrees** | **int32** | Offset that alters the unlocked position | [default to null]
**LockedPositionOffsetDegrees** | **int32** | Offset that alters the locked position | [default to null]
**DetachedCylinder** | **bool** | Flag that indicates that the inner side of the used cylinder is detached from the outer side | [optional] [default to null]
**BatteryType** | **int32** | The type of the batteries present in the smart lock: 0 .. alkali, 1 .. accumulator, 2 .. lithium | [default to null]
**AutoLock** | **bool** | New separate flag with FW &gt;&#x3D; 2.7.8/1.9.1: The Auto Lock feature automatically locks your door when it has been unlocked for a certain period of time | [optional] [default to null]
**AutoLockTimeout** | **int32** | Seconds until the smart lock relocks itself after it has been unlocked. FW &lt; 2.7.8/1.9.1: No auto relock if value is 0, FW &gt;&#x3D; 2.7.8/1.9.1: has to be &gt;&#x3D;2 (defaults to 2 for values &lt;2 if autoLock is set to true) | [optional] [default to null]
**AutoUpdateEnabled** | **bool** | Flag that indicates if available firmware updates for the deviceshould be installed automatically | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


