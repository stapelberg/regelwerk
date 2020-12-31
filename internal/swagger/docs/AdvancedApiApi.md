# \AdvancedApiApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Action**](AdvancedApiApi.md#Action) | **Post** /smartlock/{smartlockId}/action/advanced | Smartlock Action with Callback
[**PostLock**](AdvancedApiApi.md#PostLock) | **Post** /smartlock/{smartlockId}/action/lock/advanced | Lock a smartlock
[**PostLock_0**](AdvancedApiApi.md#PostLock_0) | **Post** /smartlock/{smartlockId}/action/unlock/advanced | Unlock a smartlock
[**Put**](AdvancedApiApi.md#Put) | **Put** /smartlock/auth/advanced | Creates asynchronous smartlock authorizations


# **Action**
> AdvancedConfirmationResponse Action(ctx, body, smartlockId)
Smartlock Action with Callback



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAction**](SmartlockAction.md)| Smartlock action representation | 
  **smartlockId** | **string**| The smartlock id | 

### Return type

[**AdvancedConfirmationResponse**](AdvancedConfirmationResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLock**
> AdvancedConfirmationResponse PostLock(ctx, smartlockId)
Lock a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **string**| The smartlock id | 

### Return type

[**AdvancedConfirmationResponse**](AdvancedConfirmationResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLock_0**
> AdvancedConfirmationResponse PostLock_0(ctx, smartlockId)
Unlock a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **string**| The smartlock id | 

### Return type

[**AdvancedConfirmationResponse**](AdvancedConfirmationResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put**
> AdvancedConfirmationResponse Put(ctx, body)
Creates asynchronous smartlock authorizations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlocksAuthAdvancedCreate**](SmartlocksAuthAdvancedCreate.md)| Smartlock authorization create representation | 

### Return type

[**AdvancedConfirmationResponse**](AdvancedConfirmationResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

