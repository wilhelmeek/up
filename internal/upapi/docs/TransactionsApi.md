# \TransactionsApi

All URIs are relative to *https://api.up.com.au/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAccountIdTransactionsGet**](TransactionsApi.md#AccountsAccountIdTransactionsGet) | **Get** /accounts/{accountId}/transactions | List transactions by account
[**TransactionsGet**](TransactionsApi.md#TransactionsGet) | **Get** /transactions | List transactions
[**TransactionsIdGet**](TransactionsApi.md#TransactionsIdGet) | **Get** /transactions/{id} | Retrieve transaction



## AccountsAccountIdTransactionsGet

> ListTransactionsResponse AccountsAccountIdTransactionsGet(ctx, accountId, optional)

List transactions by account

Retrieve a list of all transactions for a specific account. The returned list is [paginated](#pagination) and can be scrolled by following the `next` and `prev` links where present. To narrow the results to a specific date range pass one or both of `filter[since]` and `filter[until]` in the query string. These filter parameters **should not** be used for pagination. Results are ordered newest first to oldest last. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| The unique identifier for the account.  | 
 **optional** | ***AccountsAccountIdTransactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountsAccountIdTransactionsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of records to return in each page.  | 
 **filterSince** | **optional.Time**| The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterUntil** | **optional.Time**| The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> ListTransactionsResponse TransactionsGet(ctx, optional)

List transactions

Retrieve a list of all transactions across all accounts for the currently authenticated user. The returned list is [paginated](#pagination) and can be scrolled by following the `next` and `prev` links where present. To narrow the results to a specific date range pass one or both of `filter[since]` and `filter[until]` in the query string. These filter parameters **should not** be used for pagination. Results are ordered newest first to oldest last. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TransactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransactionsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of records to return in each page.  | 
 **filterSince** | **optional.Time**| The start date-time from which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 
 **filterUntil** | **optional.Time**| The end date-time up to which to return records, formatted according to rfc-3339. Not to be used for pagination purposes.  | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsIdGet

> GetTransactionResponse TransactionsIdGet(ctx, id)

Retrieve transaction

Retrieve a specific transaction by providing its unique identifier. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The unique identifier for the transaction.  | 

### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

