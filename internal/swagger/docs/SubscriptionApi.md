# \SubscriptionApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](SubscriptionApi.md#Get) | **Get** /subscription | Get a list of subscriptions
[**Get_0**](SubscriptionApi.md#Get_0) | **Get** /subscription/{subscriptionId} | Get a subscription


# **Get**
> []Subscription Get(ctx, )
Get a list of subscriptions



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> Subscription Get_0(ctx, subscriptionId)
Get a subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **int32**| The subscription id | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

