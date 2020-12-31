# \AccountSubscriptionApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](AccountSubscriptionApi.md#Get) | **Get** /app/account/subscription | Get a list of account subscriptions
[**Post**](AccountSubscriptionApi.md#Post) | **Post** /account/subscription/pay | Starts a payment for an account and returns a payment url
[**Post_0**](AccountSubscriptionApi.md#Post_0) | **Post** /account/subscription/{id}/activate | Activates a previously terminated subscription
[**Post_1**](AccountSubscriptionApi.md#Post_1) | **Post** /account/subscription/{id}/terminate | Terminates a running subscription


# **Get**
> []AccountSubscription Get(ctx, )
Get a list of account subscriptions



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountSubscription**](AccountSubscription.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> string Post(ctx, )
Starts a payment for an account and returns a payment url



### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, id)
Activates a previously terminated subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The account subscription unique id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_1**
> Post_1(ctx, id)
Terminates a running subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The account subscription unique id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

