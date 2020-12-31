# \ApiKeyApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ApiKeyApi.md#Delete) | **Delete** /api/key/{apiKeyId} | Delete an api key
[**Delete_0**](ApiKeyApi.md#Delete_0) | **Delete** /api/key/{apiKeyId}/advanced | Delete an advanced api key
[**Delete_1**](ApiKeyApi.md#Delete_1) | **Delete** /api/key/{apiKeyId}/token/{id} | Delete an api key token
[**Get**](ApiKeyApi.md#Get) | **Get** /api/key | Get a list of api keys
[**Get_0**](ApiKeyApi.md#Get_0) | **Get** /api/key/{apiKeyId}/advanced | Get an advanced api key
[**Get_1**](ApiKeyApi.md#Get_1) | **Get** /api/key/{apiKeyId}/token | Get a list of api key tokens
[**Post**](ApiKeyApi.md#Post) | **Post** /api/key/{apiKeyId} | Update an api key
[**Post_0**](ApiKeyApi.md#Post_0) | **Post** /api/key/{apiKeyId}/advanced | Update an advanced api key
[**Post_1**](ApiKeyApi.md#Post_1) | **Post** /api/key/{apiKeyId}/token/{id} | Update an api key token
[**Put**](ApiKeyApi.md#Put) | **Put** /api/key | Create an api key
[**Put_0**](ApiKeyApi.md#Put_0) | **Put** /api/key/{apiKeyId}/advanced | Create an advanced api key
[**Put_1**](ApiKeyApi.md#Put_1) | **Put** /api/key/{apiKeyId}/token | Create an api key token


# **Delete**
> Delete(ctx, apiKeyId)
Delete an api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| The api key id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_0**
> Delete_0(ctx, apiKeyId)
Delete an advanced api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| The api key id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_1**
> Delete_1(ctx, apiKeyId, id)
Delete an api key token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| The api key id | 
  **id** | **string**| The api key token id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> []ApiKey Get(ctx, )
Get a list of api keys



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiKey**](ApiKey.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> AdvancedApiKey Get_0(ctx, apiKeyId)
Get an advanced api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| The api key id | 

### Return type

[**AdvancedApiKey**](AdvancedApiKey.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> []ApiKeyToken Get_1(ctx, apiKeyId)
Get a list of api key tokens



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | **int32**| The api key id | 

### Return type

[**[]ApiKeyToken**](ApiKeyToken.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, body, apiKeyId)
Update an api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiKeyUpdate**](ApiKeyUpdate.md)| Api key update representation | 
  **apiKeyId** | **int32**| The api key id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, body, apiKeyId)
Update an advanced api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvancedApiKeyUpdate**](AdvancedApiKeyUpdate.md)| Update for advaced api key representation | 
  **apiKeyId** | **int32**| The api key id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_1**
> Post_1(ctx, body, apiKeyId, id)
Update an api key token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiKeyTokenUpdate**](ApiKeyTokenUpdate.md)| Api key token update representation | 
  **apiKeyId** | **int32**| The api key id | 
  **id** | **string**| The api key token id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put**
> ApiKey Put(ctx, body)
Create an api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiKeyCreate**](ApiKeyCreate.md)| Api key create representation | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put_0**
> Put_0(ctx, body, apiKeyId)
Create an advanced api key



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvancedApiKeyCreate**](AdvancedApiKeyCreate.md)| Apply for advaced api key representation | 
  **apiKeyId** | **int32**| The api key id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put_1**
> ApiKeyToken Put_1(ctx, body, apiKeyId)
Create an api key token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiKeyTokenCreate**](ApiKeyTokenCreate.md)| Api key token create representation | 
  **apiKeyId** | **int32**| The api key id | 

### Return type

[**ApiKeyToken**](ApiKeyToken.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

