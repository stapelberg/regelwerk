# SmartlockState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** | The smartlock mode: 0 .. uninitialized, 1 .. pairing, 2 .. door (default), 3 .. continuous (type&#x3D;2 only), 4 .. maintenance | [default to null]
**State** | **int32** | The smartlock state: type&#x3D;0: 0 .. uncalibrated, 1 .. locked, 2 .. unlocking, 3 .. unlocked, 4 .. locking, 5 .. unlatched, 6 .. unlocked (lock &#39;n&#39; go), 7 .. unlatching, 254 .. motor blocked, 255 .. undefined; type&#x3D;2: 0 .. untrained, 1 .. online, 3 .. ring to open active, 5 .. open, 7 .. opening, 253 .. boot run, 255 .. undefined | [default to null]
**Trigger** | **int32** |  The state trigger: 0 .. system, 1 .. manual, 2 .. button, 3 .. automatic, 4 .. web (type&#x3D;1 only), 5 .. app (type&#x3D;1 only), 6 .. continuous mode (type&#x3D;2 only), 7 .. accessory (type&#x3D;3 only) | [default to null]
**LastAction** | **int32** | The action: type&#x3D;0: 1 .. unlock, 2 .. lock, 3 .. unlatch, 4 .. lock &#39;n&#39; go, 5 .. lock &#39;n&#39; go with unlatch; type&#x3D;1: 1 .. unlock; type&#x3D;2: 1 .. activate ring to open, 2 .. deactivate ring to open, 3 .. open (electric strike actuation) | [default to null]
**BatteryCritical** | **bool** | True if the battery state of the device is critical | [default to null]
**BatteryCharging** | **bool** | True if a Nuki battery pack in a Smart Lock is currently charging | [optional] [default to null]
**BatteryCharge** | **int32** | Remaining capacity of a Nuki battery pack in % | [optional] [default to null]
**KeypadBatteryCritical** | **bool** | True if the battery of a paired Keypad is critical (only available for supported devices) | [optional] [default to null]
**DoorState** | **int32** | The door state: 0 .. unavailable, 1 .. deactivated, 2 .. door closed, 3 .. door opened, 4 .. door state unknown, 5 .. calibrating | [default to null]
**RingToOpenTimer** | **int32** | [deprecated] Remaining ring to open time; 0 if ring to open is not active (type&#x3D;2 only) | [default to null]
**RingToOpenEnd** | [**time.Time**](time.Time.md) | End date of ring to open timeout; null if ring to open is not active (type&#x3D;2 only) | [optional] [default to null]
**NightMode** | **bool** | True if night mode currently active | [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


