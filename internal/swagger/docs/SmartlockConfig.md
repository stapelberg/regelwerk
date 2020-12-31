# SmartlockConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the smartlock for new users | [default to null]
**Latitude** | **float32** | The latitude of the smartlock position | [default to null]
**Longitude** | **float32** | The longitude of the smartlock position | [default to null]
**Capabilities** | **int32** | The capabilities indicate whether door opening via app is possible, RTO is possible or both: 0 .. only door opening possible, 1 .. both possible, 2 .. only RTO possible (only for type&#x3D;2) | [optional] [default to null]
**AutoUnlatch** | **bool** | True if the door should be unlatched on unlocking (knob) (only for type&#x3D;1 and type&#x3D;3) | [optional] [default to null]
**PairingEnabled** | **bool** | True if the pairing is allowed via the smartlock button | [optional] [default to null]
**ButtonEnabled** | **bool** | True if the button on the smartlock is enabled | [optional] [default to null]
**LedEnabled** | **bool** | True if the LED on the smartlock is enabled | [optional] [default to null]
**LedBrightness** | **int32** | The brightness of the LED: 0 .. off, 5 .. max (only for type&#x3D;1 and type&#x3D;3) | [optional] [default to null]
**TimezoneOffset** | **int32** | [deprecated] The timezone offset (in minutes) | [default to null]
**DaylightSavingMode** | **int32** | [deprecated] The daylight saving mode: 0 .. off, 1 .. european | [optional] [default to null]
**FobPaired** | **bool** | True if a fob is paired with the smartlock | [optional] [default to null]
**FobAction1** | **int32** | The fob action if button is pressed once: type&#x3D;0: 0 .. none, 1 .. unlock, 2 .. lock, 3 .. lock &#39;n&#39; go, 4 .. intelligent (lock/unlocked based on the current state); type&#x3D;2: 0 .. none, 1 .. toggle ring to open, 2 .. activate ring to open, 3 .. deactivate ring to open, 7 .. open (electric strike actuation), 8 .. ring | [optional] [default to null]
**FobAction2** | **int32** | The fob action if button is pressed twice: type&#x3D;0: 0 .. none, 1 .. unlock, 2 .. lock, 3 .. lock &#39;n&#39; go, 4 .. intelligent (lock/unlocked based on the current state); type&#x3D;2: 0 .. none, 1 .. toggle ring to open, 2 .. activate ring to open, 3 .. deactivate ring to open, 7 .. open (electric strike actuation), 8 .. ring | [optional] [default to null]
**FobAction3** | **int32** | The fob action if button is pressed 3 times: type&#x3D;0: 0 .. none, 1 .. unlock, 2 .. lock, 3 .. lock &#39;n&#39; go, 4 .. intelligent (lock/unlocked based on the current state); type&#x3D;2: 0 .. none, 1 .. toggle ring to open, 2 .. activate ring to open, 3 .. deactivate ring to open, 7 .. open (electric strike actuation), 8 .. ring | [optional] [default to null]
**SingleLock** | **bool** | True if the smartlock should only lock once (instead of twice) (only for type&#x3D;1) | [default to null]
**OperatingMode** | **int32** | The operating mode of the opener (only for type&#x3D;2): 0x00 .. generic door opener, 0x01 .. analogue intercom, 0x02 .. digital intercom, 0x03 .. digital intercom Siedle, 0x04 .. digital intercom TCS, 0x05 .. digital intercom Bticino, 0x06 .. analog intercom Siedle HTS, 0x07 .. digital intercom STR, 0x08 .. digital intercom Ritto, 0x09 .. digital intercom Fermax, 0x0A .. digital intercom Comelit, 0x0B .. digital intercom Urmet BiBus, 0x0C .. digital intercom Urmet 2Voice, 0x0D .. digital intercom Golmar, 0x0E .. digital intercom SKS, 0x0F .. digital intercom Spare | [optional] [default to null]
**AdvertisingMode** | **int32** | The advertising mode (battery saving): 0 .. automatic, 1 .. normal, 2 .. slow, 3 .. slowest | [default to null]
**KeypadPaired** | **bool** | True if a keypad is paired with the smartlock | [optional] [default to null]
**HomekitState** | **int32** | The homekit state: 0 .. unavailable, 1 .. disabled, 2 .. enabled, 3 .. enabled &amp; paired | [optional] [default to null]
**TimezoneId** | **int32** | The timezone id (check https://developer.nuki.io for ids) | [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


