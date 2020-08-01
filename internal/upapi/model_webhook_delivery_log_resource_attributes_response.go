/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package upapi
// WebhookDeliveryLogResourceAttributesResponse Information about the response that was received from the webhook URL. 
type WebhookDeliveryLogResourceAttributesResponse struct {
	// The HTTP status code received in the response. 
	StatusCode int32 `json:"statusCode"`
	// The payload that was recieved in the response body. 
	Body string `json:"body"`
}