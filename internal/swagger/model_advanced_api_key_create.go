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

type AdvancedApiKeyCreate struct {
	// The name of the company for which you apply for access
	Name string `json:"name"`
	// The country of the headquarter or the country where you are mainly doing business in
	Country string `json:"country"`
	// Describe the services and/or products you offer to your customers and how your customers would use Nuki devices in their processes
	Description string `json:"description"`
	// The application type
	Type_ string `json:"type"`
	// A website where we can find more information on the company and its business model
	Url string `json:"url"`
	// An email address where we can contact you for checks on your application
	Email string `json:"email"`
	// The URL where our webhooks should point to
	WebhookUrl string `json:"webhookUrl"`
	// The features to trigger webhooks, for all types except 'ONLY_SECRET'
	WebhookFeatures []string `json:"webhookFeatures"`
}
