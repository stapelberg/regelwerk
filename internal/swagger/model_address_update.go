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

type AddressUpdate struct {
	// The smartlocks for this address
	SmartlockIds []int64 `json:"smartlockIds"`
	// The optional settings
	Settings map[string]interface{} `json:"settings,omitempty"`
}
