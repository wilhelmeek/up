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
// WebhookDeliveryLogResourceAttributes struct for WebhookDeliveryLogResourceAttributes
type WebhookDeliveryLogResourceAttributes struct {
	Request WebhookDeliveryLogResourceAttributesRequest `json:"request"`
	Response *WebhookDeliveryLogResourceAttributesResponse `json:"response"`
	// The success or failure status of this delivery attempt. 
	DeliveryStatus WebhookDeliveryStatusEnum `json:"deliveryStatus"`
	// The date-time at which this log entry was created. 
	CreatedAt time.Time `json:"createdAt"`
}
