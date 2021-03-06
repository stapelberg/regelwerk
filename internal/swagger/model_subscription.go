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

type Subscription struct {
	// The unique subscription id
	SubscriptionId int32 `json:"subscriptionId"`
	// The title per language
	Titles map[string]string `json:"titles"`
	// The description per language
	Descriptions map[string]string `json:"descriptions"`
	// The quantity of authorizations
	Quantity int32 `json:"quantity"`
	// The length with number and unit (e.g. 30d, 2y)
	Length string `json:"length"`
	// The price
	Price float64 `json:"price"`
	// The optional gift article
	GiftArticle string `json:"giftArticle,omitempty"`
	// True if the subscription is available for new users
	Available bool `json:"available"`
	// The update date
	UpdateDate int64 `json:"updateDate,omitempty"`
}
