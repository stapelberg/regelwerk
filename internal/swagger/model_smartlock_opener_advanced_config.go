/*
 * Nuki API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Contact: contact@nuki.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SmartlockOpenerAdvancedConfig struct {
	// The database ID of the connected intercom
	IntercomId int32 `json:"intercomId"`
	// Method to switch between data and analogue mode
	BusModeSwitch int32 `json:"busModeSwitch"`
	// Duration of the short circuit for BUS mode switching in ms
	ShortCircuitDuration int32 `json:"shortCircuitDuration"`
	// Delay of electric strike activation in ms after lock action 3 'electric strike actuation'
	ElectricStrikeDelay int32 `json:"electricStrikeDelay"`
	// Random electricStrikeDelay (range 3000 - 7000 ms) in order to simulate a person inside actuating the electric strike
	RandomElectricStrikeDelay bool `json:"randomElectricStrikeDelay"`
	// Duration in ms of electric strike actuation lock action 3 'electric strike actuation'
	ElectricStrikeDuration int32 `json:"electricStrikeDuration"`
	// Flag to disable RTO after ring
	DisableRtoAfterRing bool `json:"disableRtoAfterRing"`
	// After this period of time in minutes, RTO gets deactivated automatically
	RtoTimeout int32 `json:"rtoTimeout"`
	// The doorbell supression bitmask: first bit (least significant) .. whenever the doorbell rings and CM and RTO are inactive, second bit .. RTO is active, third bit .. CM is active
	DoorbellSuppression int32 `json:"doorbellSuppression"`
	// Duration in ms of doorbell suppression (only in Operating mode 2 'digital Intercom')
	DoorbellSuppressionDuration int32 `json:"doorbellSuppressionDuration"`
	// The sound for ring: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3
	SoundRing int32 `json:"soundRing"`
	// The sound for open: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3
	SoundOpen int32 `json:"soundOpen"`
	// The sound for RTO: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3
	SoundRto int32 `json:"soundRto"`
	// The sound for CM: 0 .. no sound, 1 .. Sound1, 2 .. Sound2, 3 .. Sound3
	SoundCm int32 `json:"soundCm"`
	// The sound confirmation: 0 .. no sound, 1 .. sound
	SoundConfirmation int32 `json:"soundConfirmation"`
	// The sound level
	SoundLevel int32 `json:"soundLevel"`
	// The desired action, if the button is pressed once: 0 .. no action, 1 .. toggle RTO, 2 .. activate RTO, 3 .. deactivate RTO, 4 .. toggle CM, 5 .. activate CM, 6 .. deactivate CM, 7 .. open
	SingleButtonPressAction int32 `json:"singleButtonPressAction"`
	// The desired action, if the button is pressed twice: 0 .. no action, 1 .. toggle RTO, 2 .. activate RTO, 3 .. deactivate RTO, 4 .. toggle CM, 5 .. activate CM, 6 .. deactivate CM, 7 .. open
	DoubleButtonPressAction int32 `json:"doubleButtonPressAction"`
	// The type of the batteries present in the smart lock: 0 .. alkali, 1 .. accumulator, 2 .. lithium, 3 .. fixed
	BatteryType int32 `json:"batteryType"`
	// Flag that indicates if the automatic detection of the battery type is enabled
	AutomaticBatteryTypeDetection bool `json:"automaticBatteryTypeDetection,omitempty"`
	// Flag that indicates if available firmware updates for the deviceshould be installed automatically
	AutoUpdateEnabled bool `json:"autoUpdateEnabled,omitempty"`
	// The operation id - if set it's locked for another operation
	OperationId string `json:"operationId,omitempty"`
}
