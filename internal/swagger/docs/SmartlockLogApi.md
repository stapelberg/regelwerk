# \SmartlockLogApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](SmartlockLogApi.md#Get) | **Get** /smartlock/log | Get a list of smartlock logs for all of your smartlocks
[**Get_0**](SmartlockLogApi.md#Get_0) | **Get** /smartlock/{smartlockId}/log | Get a list of smartlock logs


# **Get**
> []SmartlockLog Get(ctx, optional)
Get a list of smartlock logs for all of your smartlocks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartlockLogApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartlockLogApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUserId** | **optional.Int32**| Filter for account users | 
 **fromDate** | **optional.String**| Filter for date (RFC3339) | 
 **toDate** | **optional.String**| Filter for date (RFC3339) | 
 **action** | **optional.Int32**| Filter for action | 
 **id** | **optional.Int32**| Filter for older logs | 
 **limit** | **optional.Int32**| Amount of logs (max: 50) | [default to 20]

### Return type

[**[]SmartlockLog**](SmartlockLog.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []SmartlockLog Get_0(ctx, smartlockId, optional)
Get a list of smartlock logs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smartlockId** | **int32**| The smartlock id | 
 **optional** | ***SmartlockLogApiGet_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartlockLogApiGet_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authId** | **optional.String**| Filter for auths | 
 **accountUserId** | **optional.Int32**| Filter for account users | 
 **fromDate** | **optional.String**| Filter for date (RFC3339) | 
 **toDate** | **optional.String**| Filter for date (RFC3339) | 
 **action** | **optional.Int32**| Filter for action | 
 **id** | **optional.String**| Filter for older logs | 
 **limit** | **optional.Int32**| Amount of logs (max: 50) | [default to 20]

### Return type

[**[]SmartlockLog**](SmartlockLog.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

