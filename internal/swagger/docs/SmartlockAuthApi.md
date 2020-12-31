# \SmartlockAuthApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](SmartlockAuthApi.md#Delete) | **Delete** /smartlock/auth | Deletes smartlock authorizations asynchronously
[**Delete_0**](SmartlockAuthApi.md#Delete_0) | **Delete** /smartlock/{smartlockId}/auth/{id} | Deletes asynchronous a smartlock authorization
[**Get**](SmartlockAuthApi.md#Get) | **Get** /smartlock/auth | Get a list of smartlock authorizations for your smartlocks
[**Get_0**](SmartlockAuthApi.md#Get_0) | **Get** /smartlock/{smartlockId}/auth | Get a list of smartlock authorizations
[**Get_1**](SmartlockAuthApi.md#Get_1) | **Get** /smartlock/{smartlockId}/auth/{id} | Get a smartlock authorization
[**Post**](SmartlockAuthApi.md#Post) | **Post** /smartlock/auth | Updates smartlock authorizations asynchronously
[**Post_0**](SmartlockAuthApi.md#Post_0) | **Post** /smartlock/{smartlockId}/auth/{id} | Updates asynchronous a smartlock authorization
[**Put**](SmartlockAuthApi.md#Put) | **Put** /smartlock/auth | Creates asynchronous smartlock authorizations
[**Put_0**](SmartlockAuthApi.md#Put_0) | **Put** /smartlock/{smartlockId}/auth | Creates asynchronous a smartlock authorization


# **Delete**
> Delete(ctx, body)
Deletes smartlock authorizations asynchronously



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | **[]string**| Smartlock authorization IDS to delete | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_0**
> Delete_0(ctx, smartlockId, id)
Deletes asynchronous a smartlock authorization



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 
  **id** | **string**| The smartlock authorization unique id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> []SmartlockAuth Get(ctx, optional)
Get a list of smartlock authorizations for your smartlocks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartlockAuthApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartlockAuthApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUserId** | **optional.Int32**| Filter for account users:  set to a positive number will filter for authorizations with this specific accountUserId, set to a negative number will filter without set accountUserId | 
 **types** | **optional.String**| Filter for authorization&#39;s types (comma-separated eg: 0,2,3) | 

### Return type

[**[]SmartlockAuth**](SmartlockAuth.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []SmartlockAuth Get_0(ctx, smartlockId, optional)
Get a list of smartlock authorizations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 
 **optional** | ***SmartlockAuthApiGet_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartlockAuthApiGet_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **types** | **optional.String**| Filter for smartlock authorization&#39;s types (comma-separated eg: 0,2,3) | 

### Return type

[**[]SmartlockAuth**](SmartlockAuth.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> SmartlockAuth Get_1(ctx, smartlockId, id)
Get a smartlock authorization



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 
  **id** | **string**| The smartlock auth unique id | 

### Return type

[**SmartlockAuth**](SmartlockAuth.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, body)
Updates smartlock authorizations asynchronously



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]SmartlockAuthMultiUpdate**](SmartlockAuthMultiUpdate.md)| Smartlock authorization update representations | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, body, smartlockId, id)
Updates asynchronous a smartlock authorization



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAuthUpdate**](SmartlockAuthUpdate.md)| Smartlock authorization update representation | 
  **smartlockId** | **int32**| The smartlock id | 
  **id** | **string**| The smartlock authorization unique id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put**
> Put(ctx, body)
Creates asynchronous smartlock authorizations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlocksAuthCreate**](SmartlocksAuthCreate.md)| Smartlock authorization create representation | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put_0**
> Put_0(ctx, body, smartlockId)
Creates asynchronous a smartlock authorization



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAuthCreate**](SmartlockAuthCreate.md)| Smartlock authorization create representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

