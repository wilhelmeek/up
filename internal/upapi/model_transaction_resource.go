/*
 * Up API
 *
 * The Up API gives you programmatic access to your balances and transaction data. You can request past transactions or set up webhooks to receive real-time events when new transactions hit your account. It’s new, it’s exciting and it’s just the beginning. 
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package upapi
// TransactionResource struct for TransactionResource
type TransactionResource struct {
	// The unique identifier for this transaction. 
	Id string `json:"id"`
	// The type of this resource: `transactions`
	Type string `json:"type"`
	Attributes TransactionResourceAttributes `json:"attributes"`
	Links AccountResourceLinks `json:"links,omitempty"`
	Relationships TransactionResourceRelationships `json:"relationships"`
}
