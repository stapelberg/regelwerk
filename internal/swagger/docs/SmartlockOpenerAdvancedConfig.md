# SmartlockOpenerAdvancedConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntercomId** | **int32** | The database ID of the connected intercom | [default to null]
**BusModeSwitch** | **int32** | Method to switch between data and analogue mode | [default to null]
**ShortCircuitDuration** | **int32** | Duration of the short circuit for BUS mode switching in ms | [default to null]
**ElectricStrikeDelay** | **int32** | Delay of electric strike activation in ms after lock action 3 &#39;electric strike actuation&#39; | [default to null]
**RandomElectricStrikeDelay** | **bool** | Random electricStrikeDelay (range 3000 - 7000 ms) in order to simulate a person inside actuating the electric strike | [default to null]
**ElectricStrikeDuration** | **int32** | Duration in ms of electric strike actuation lock action 3 &#39;electric strike actuation&#39; | [default to null]
**DisableRtoAfterRing** | **bool** | Flag to disable RTO after ring | [default to null]
**RtoTimeout** | **int32** | After this period of time in minutes, RTO gets deactivated automatically | [default to null]
**DoorbellSuppression** | **int32** | The doorbell supression bitmask: first bit (least significant) .. whenever the doorbell rings and CM and RTO are inactive, second bit .. RTO is active, third bit .. CM is active | [default to null]
**DoorbellSuppressionDuration** | **int32** | Duration in ms of doorbell suppression (only in Operating mode 2 &#39;digital Intercom&#39;) | [default to null]
**SoundRing** | **int32** | The sound for ring: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3 | [default to null]
**SoundOpen** | **int32** | The sound for open: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3 | [default to null]
**SoundRto** | **int32** | The sound for RTO: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3 | [default to null]
**SoundCm** | **int32** | The sound for CM: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3 | [default to null]
**SoundConfirmation** | **int32** | The sound confirmation: 0 .. no sound, 1 .. sound | [default to null]
**SoundLevel** | **int32** | The sound level | [default to null]
**SingleButtonPressAction** | **int32** | The desired action, if the button is pressed once: 0 .. no action, 1 .. toggle RTO, 2 .. activate RTO, 3 .. deactivate RTO, 4 .. toggle CM, 5 .. activate CM, 6 .. deactivate CM, 7 .. open | [default to null]
**DoubleButtonPressAction** | **int32** | The desired action, if the button is pressed twice: 0 .. no action, 1 .. toggle RTO, 2 .. activate RTO, 3 .. deactivate RTO, 4 .. toggle CM, 5 .. activate CM, 6 .. deactivate CM, 7 .. open | [default to null]
**BatteryType** | **int32** | The type of the batteries present in the smart lock: 0 .. alkali, 1 .. accumulator, 2 .. lithium, 3 .. fixed | [default to null]
**AutomaticBatteryTypeDetection** | **bool** | Flag that indicates if the automatic detection of the battery type is enabled | [optional] [default to null]
**AutoUpdateEnabled** | **bool** | Flag that indicates if available firmware updates for the deviceshould be installed automatically | [optional] [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


