# \SmartlockApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](SmartlockApi.md#Delete) | **Delete** /smartlock/{smartlockId} | Delete a smartlock
[**Get**](SmartlockApi.md#Get) | **Get** /smartlock | Get a list of smartlocks
[**Get_0**](SmartlockApi.md#Get_0) | **Get** /smartlock/{smartlockId} | Get a smartlock
[**Post**](SmartlockApi.md#Post) | **Post** /smartlock/{smartlockId} | Update a smartlock
[**PostLock**](SmartlockApi.md#PostLock) | **Post** /smartlock/{smartlockId}/action/lock | Lock a smartlock
[**PostUnlock**](SmartlockApi.md#PostUnlock) | **Post** /smartlock/{smartlockId}/action/unlock | Unlock a smartlock
[**Post_0**](SmartlockApi.md#Post_0) | **Post** /smartlock/{smartlockId}/action | Lock &amp; unlock a smartlock with options
[**Post_1**](SmartlockApi.md#Post_1) | **Post** /smartlock/{smartlockId}/admin/pin | Updates a smartlock admin pin
[**Post_2**](SmartlockApi.md#Post_2) | **Post** /smartlock/{smartlockId}/advanced/config | Updates a smartlock advanced config
[**Post_3**](SmartlockApi.md#Post_3) | **Post** /smartlock/{smartlockId}/advanced/openerconfig | Updates an opener advanced config
[**Post_4**](SmartlockApi.md#Post_4) | **Post** /smartlock/{smartlockId}/advanced/smartdoorconfig | Updates a smartdoor advanced config
[**Post_5**](SmartlockApi.md#Post_5) | **Post** /smartlock/{smartlockId}/config | Updates a smartlock config
[**Post_6**](SmartlockApi.md#Post_6) | **Post** /smartlock/{smartlockId}/sync | Syncs a smartlock
[**Post_7**](SmartlockApi.md#Post_7) | **Post** /smartlock/{smartlockId}/web/config | Updates a smartlock web config
[**Put**](SmartlockApi.md#Put) | **Put** /smartlock | Create a smartlock


# **Delete**
> Delete(ctx, smartlockId)
Delete a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> []Smartlock Get(ctx, optional)
Get a list of smartlocks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartlockApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartlockApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authId** | **optional.Int32**| Filter for authId | 
 **type_** | **optional.Int32**| Filter for type | 

### Return type

[**[]Smartlock**](Smartlock.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> Smartlock Get_0(ctx, smartlockId)
Get a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 

### Return type

[**Smartlock**](Smartlock.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, body, smartlockId)
Update a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockUpdate**](SmartlockUpdate.md)| Smartlock update representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLock**
> PostLock(ctx, smartlockId)
Lock a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **string**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUnlock**
> PostUnlock(ctx, smartlockId)
Unlock a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **string**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, body, smartlockId)
Lock & unlock a smartlock with options



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAction**](SmartlockAction.md)| Smartlock action representation | 
  **smartlockId** | **string**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_1**
> Post_1(ctx, body, smartlockId)
Updates a smartlock admin pin



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAdminPinUpdate**](SmartlockAdminPinUpdate.md)| Smartlock admin pin update representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_2**
> Post_2(ctx, body, smartlockId)
Updates a smartlock advanced config



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockAdvancedConfig**](SmartlockAdvancedConfig.md)| Smartlock config update representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_3**
> Post_3(ctx, body, smartlockId)
Updates an opener advanced config



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockOpenerAdvancedConfig**](SmartlockOpenerAdvancedConfig.md)| Opener advanced config update representation | 
  **smartlockId** | **int32**| The smartlock (opener) id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_4**
> Post_4(ctx, body, smartlockId)
Updates a smartdoor advanced config



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockSmartdoorAdvancedConfig**](SmartlockSmartdoorAdvancedConfig.md)| Smartdoor advanced config update representation | 
  **smartlockId** | **int32**| The smartdoor id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_5**
> Post_5(ctx, body, smartlockId)
Updates a smartlock config



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockConfig**](SmartlockConfig.md)| Smartlock config update representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_6**
> Post_6(ctx, smartlockId)
Syncs a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **string**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_7**
> Post_7(ctx, body, smartlockId)
Updates a smartlock web config



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockWebConfig**](SmartlockWebConfig.md)| Smartlock web config update representation | 
  **smartlockId** | **int32**| The smartlock id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put**
> Smartlock Put(ctx, body)
Create a smartlock



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SmartlockCreate**](SmartlockCreate.md)| Smartlock create representation | 

### Return type

[**Smartlock**](Smartlock.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

