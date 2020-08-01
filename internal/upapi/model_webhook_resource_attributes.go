/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package upapi
import (
	"time"
)
// WebhookResourceAttributes struct for WebhookResourceAttributes
type WebhookResourceAttributes struct {
	// The URL that this webhook is configured to `POST` events to. 
	Url string `json:"url"`
	// An optional description that was provided at the time the webhook was created. 
	Description *string `json:"description"`
	// A shared secret key used to sign all webhook events sent to the configured webhook URL. This field is returned only once, upon the initial creation of the webhook. If lost, create a new webhook and delete this webhook.  The webhook URL receives a request with a `X-Up-Authenticity-Signature` header, which is the SHA-256 HMAC of the entire raw request body signed using this `secretKey`. It is advised to compute and check this signature to verify the authenticity of requests sent to the webhook URL. See [Handling webhook events](#callback_post_webhookURL) for full details. 
	SecretKey string `json:"secretKey,omitempty"`
	// The date-time at which this webhook was created. 
	CreatedAt time.Time `json:"createdAt"`
}