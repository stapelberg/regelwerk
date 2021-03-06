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

type AccountSubscription struct {
	// The id
	Id *ObjectId `json:"id"`
	// The account id
	AccountId int32 `json:"accountId"`
	// The subscription id
	SubscriptionId int32 `json:"subscriptionId"`
	// The actual period
	Period int32 `json:"period"`
	// The payment type
	PaymentType string `json:"paymentType"`
	// The quantity of authorizations
	Quantity int32 `json:"quantity"`
	// The optional gift article after first purchase
	GiftArticle string `json:"giftArticle,omitempty"`
	// The status
	Status string `json:"status"`
	// The start date
	StartDate int64 `json:"startDate,omitempty"`
	// The period end date
	PeriodEndDate int64 `json:"periodEndDate,omitempty"`
	// The next payment date
	NextPaymentDate int64 `json:"nextPaymentDate,omitempty"`
}
