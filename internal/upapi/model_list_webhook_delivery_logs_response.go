/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package upapi
// ListWebhookDeliveryLogsResponse Successful response to get all delivery logs for a webhook. This returns a [paginated](#pagination) list of delivery logs, which can be scrolled by following the `next` and `prev` links if present. 
type ListWebhookDeliveryLogsResponse struct {
	// The list of delivery logs returned in this response. 
	Data []WebhookDeliveryLogResource `json:"data"`
	Links ListAccountsResponseLinks `json:"links"`
}
