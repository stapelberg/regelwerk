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

type ApiKeyTokenUpdate struct {
	// The description
	Description string `json:"description,omitempty"`
	// The list of scopes
	Scopes []string `json:"scopes,omitempty"`
}
