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

import (
	"time"
)

type SmartlockState struct {
	// The smartlock mode: 0 .. uninitialized, 1 .. pairing, 2 .. door (default), 3 .. continuous (type=2 only), 4 .. maintenance
	Mode int32 `json:"mode"`
	// The smartlock state: type=0: 0 .. uncalibrated, 1 .. locked, 2 .. unlocking, 3 .. unlocked, 4 .. locking, 5 .. unlatched, 6 .. unlocked (lock 'n' go), 7 .. unlatching, 254 .. motor blocked, 255 .. undefined; type=2: 0 .. untrained, 1 .. online, 3 .. ring to open active, 5 .. open, 7 .. opening, 253 .. boot run, 255 .. undefined
	State int32 `json:"state"`
	//  The state trigger: 0 .. system, 1 .. manual, 2 .. button, 3 .. automatic, 4 .. web (type=1 only), 5 .. app (type=1 only), 6 .. continuous mode (type=2 only), 7 .. accessory (type=3 only)
	Trigger int32 `json:"trigger"`
	// The action: type=0: 1 .. unlock, 2 .. lock, 3 .. unlatch, 4 .. lock 'n' go, 5 .. lock 'n' go with unlatch; type=1: 1 .. unlock; type=2: 1 .. activate ring to open, 2 .. deactivate ring to open, 3 .. open (electric strike actuation)
	LastAction int32 `json:"lastAction"`
	// True if the battery state of the device is critical
	BatteryCritical bool `json:"batteryCritical"`
	// True if a Nuki battery pack in a Smart Lock is currently charging
	BatteryCharging bool `json:"batteryCharging,omitempty"`
	// Remaining capacity of a Nuki battery pack in %
	BatteryCharge int32 `json:"batteryCharge,omitempty"`
	// True if the battery of a paired Keypad is critical (only available for supported devices)
	KeypadBatteryCritical bool `json:"keypadBatteryCritical,omitempty"`
	// The door state: 0 .. unavailable, 1 .. deactivated, 2 .. door closed, 3 .. door opened, 4 .. door state unknown, 5 .. calibrating
	DoorState int32 `json:"doorState"`
	// [deprecated] Remaining ring to open time; 0 if ring to open is not active (type=2 only)
	RingToOpenTimer int32 `json:"ringToOpenTimer"`
	// End date of ring to open timeout; null if ring to open is not active (type=2 only)
	RingToOpenEnd time.Time `json:"ringToOpenEnd,omitempty"`
	// True if night mode currently active
	NightMode bool `json:"nightMode"`
	// The operation id - if set it's locked for another operation
	OperationId string `json:"operationId,omitempty"`
}
