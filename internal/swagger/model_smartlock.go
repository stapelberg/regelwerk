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

type Smartlock struct {
	// The smartlock id
	SmartlockId int64 `json:"smartlockId"`
	// The account id
	AccountId int32 `json:"accountId"`
	// The type: 0 .. keyturner, 1 .. box, 2 .. opener, 3 .. smartdoor
	Type_ int32 `json:"type"`
	// The authorization id
	AuthId int32 `json:"authId"`
	// The name of the smartlock
	Name string `json:"name"`
	// The favorite flag
	Favorite bool `json:"favorite"`
	// The config
	Config *SmartlockConfig `json:"config,omitempty"`
	// The advanced config
	AdvancedConfig *SmartlockAdvancedConfig `json:"advancedConfig,omitempty"`
	// The opener advanced config
	OpenerAdvancedConfig *SmartlockOpenerAdvancedConfig `json:"openerAdvancedConfig,omitempty"`
	// The smartdoor advanced config
	SmartdoorAdvancedConfig *SmartlockSmartdoorAdvancedConfig `json:"smartdoorAdvancedConfig,omitempty"`
	// The web config
	WebConfig *SmartlockWebConfig `json:"webConfig,omitempty"`
	// The state
	State *SmartlockState `json:"state,omitempty"`
	// The firmware version
	FirmwareVersion int32 `json:"firmwareVersion,omitempty"`
	// The hardware version
	HardwareVersion int32 `json:"hardwareVersion,omitempty"`
	// The operation id - if set it's locked for another operation
	OperationId string `json:"operationId,omitempty"`
	// The server state: 0 .. ok, 1 .. unregistered, 2 .. auth uuid invalid, 3 .. auth invalid, 4 .. offline
	ServerState int32 `json:"serverState"`
	// The admin pin state: 0 .. ok, 1 .. missing, 2 .. invalid
	AdminPinState int32 `json:"adminPinState"`
	// The flag indicating a virtual Smart Lock
	VirtualDevice bool `json:"virtualDevice,omitempty"`
	// The creation date
	CreationDate time.Time `json:"creationDate,omitempty"`
	// The update date
	UpdateDate time.Time `json:"updateDate,omitempty"`
	Opener bool `json:"opener,omitempty"`
	Box bool `json:"box,omitempty"`
	SmartDoor bool `json:"smartDoor,omitempty"`
	Keyturner bool `json:"keyturner,omitempty"`
}
