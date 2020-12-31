# Smartlock

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmartlockId** | **int64** | The smartlock id | [default to null]
**AccountId** | **int32** | The account id | [default to null]
**Type_** | **int32** | The type: 0 .. keyturner, 1 .. box, 2 .. opener, 3 .. smartdoor | [default to null]
**AuthId** | **int32** | The authorization id | [default to null]
**Name** | **string** | The name of the smartlock | [default to null]
**Favorite** | **bool** | The favorite flag | [default to null]
**Config** | [***SmartlockConfig**](Smartlock.Config.md) | The config | [optional] [default to null]
**AdvancedConfig** | [***SmartlockAdvancedConfig**](Smartlock.AdvancedConfig.md) | The advanced config | [optional] [default to null]
**OpenerAdvancedConfig** | [***SmartlockOpenerAdvancedConfig**](Smartlock.OpenerAdvancedConfig.md) | The opener advanced config | [optional] [default to null]
**SmartdoorAdvancedConfig** | [***SmartlockSmartdoorAdvancedConfig**](Smartlock.SmartdoorAdvancedConfig.md) | The smartdoor advanced config | [optional] [default to null]
**WebConfig** | [***SmartlockWebConfig**](Smartlock.WebConfig.md) | The web config | [optional] [default to null]
**State** | [***SmartlockState**](Smartlock.State.md) | The state | [optional] [default to null]
**FirmwareVersion** | **int32** | The firmware version | [optional] [default to null]
**HardwareVersion** | **int32** | The hardware version | [optional] [default to null]
**OperationId** | **string** | The operation id - if set it&#39;s locked for another operation | [optional] [default to null]
**ServerState** | **int32** | The server state: 0 .. ok, 1 .. unregistered, 2 .. auth uuid invalid, 3 .. auth invalid, 4 .. offline | [default to null]
**AdminPinState** | **int32** | The admin pin state: 0 .. ok, 1 .. missing, 2 .. invalid | [default to null]
**VirtualDevice** | **bool** | The flag indicating a virtual Smart Lock | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The creation date | [optional] [default to null]
**UpdateDate** | [**time.Time**](time.Time.md) | The update date | [optional] [default to null]
**Opener** | **bool** |  | [optional] [default to null]
**Box** | **bool** |  | [optional] [default to null]
**SmartDoor** | **bool** |  | [optional] [default to null]
**Keyturner** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


