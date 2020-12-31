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

type SmartlockUpdate struct {
	// The new name of the smartlock
	Name string `json:"name,omitempty"`
	// True if the smartlock is favorite
	Favorite bool `json:"favorite,omitempty"`
}
