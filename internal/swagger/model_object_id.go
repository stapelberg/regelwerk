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

type ObjectId struct {
	Timestamp         int32     `json:"timestamp,omitempty"`
	MachineIdentifier int32     `json:"machineIdentifier,omitempty"`
	ProcessIdentifier int32     `json:"processIdentifier,omitempty"`
	Counter           int32     `json:"counter,omitempty"`
	Time              int64     `json:"time,omitempty"`
	Date              time.Time `json:"date,omitempty"`
	TimeSecond        int32     `json:"timeSecond,omitempty"`
}
