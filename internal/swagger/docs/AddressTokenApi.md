# \AddressTokenApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](AddressTokenApi.md#Get) | **Get** /address/token/{id} | Gives some info about address token
[**Get_0**](AddressTokenApi.md#Get_0) | **Get** /address/token/{id}/redeem | Gives an redeemed address token
[**Get_1**](AddressTokenApi.md#Get_1) | **Get** /address/{addressId}/token | Get a list of address tokens
[**Post**](AddressTokenApi.md#Post) | **Post** /address/token/{id}/redeem | Redeems an address token


# **Get**
> AddressTokenInfo Get(ctx, id)
Gives some info about address token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The token id | 

### Return type

[**AddressTokenInfo**](AddressTokenInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> AddressToken Get_0(ctx, id)
Gives an redeemed address token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The token id | 

### Return type

[**AddressToken**](AddressToken.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> []AddressToken Get_1(ctx, addressId)
Get a list of address tokens



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressId** | **int32**| The address id | 

### Return type

[**[]AddressToken**](AddressToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, id, optional)
Redeems an address token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The token id | 
 **optional** | ***AddressTokenApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressTokenApiPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.Bool**| If false no email will be send | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

