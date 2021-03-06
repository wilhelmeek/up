/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package upapi
// WebhookEventResource Provides the event data used in asynchronous webhook event callbacks to subscribed endpoints. Webhooks events have defined `eventType`s and may optionally relate to other resources within the Up API. 
type WebhookEventResource struct {
	// The unique identifier for this event. This will remain constant across delivery retries. 
	Id string `json:"id"`
	// The type of this resource: `webhook-events`
	Type string `json:"type"`
	Attributes WebhookEventResourceAttributes `json:"attributes"`
	Links AccountResourceLinks `json:"links,omitempty"`
	Relationships WebhookEventResourceRelationships `json:"relationships"`
}
