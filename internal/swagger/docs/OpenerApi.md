# \OpenerApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OpenerApi.md#Get) | **Get** /opener/brand | Get all intercom brands
[**Get_0**](OpenerApi.md#Get_0) | **Get** /opener/brand/{brandId} | Get an intercom brand
[**Get_1**](OpenerApi.md#Get_1) | **Get** /opener/intercom | Get a list of intercom models
[**Get_2**](OpenerApi.md#Get_2) | **Get** /opener/intercom/{intercomId} | Get an intercom model


# **Get**
> []OpenerIntercomBrand Get(ctx, )
Get all intercom brands



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OpenerIntercomBrand**](OpenerIntercomBrand.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> OpenerIntercomBrand Get_0(ctx, brandId)
Get an intercom brand



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brandId** | **int32**| The brand ID | 

### Return type

[**OpenerIntercomBrand**](OpenerIntercomBrand.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> []OpenerIntercomModel Get_1(ctx, optional)
Get a list of intercom models



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenerApiGet_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenerApiGet_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandId** | **optional.Int32**| Filter for brandId. Required if &#39;recentlyChanged&#39; is not set | 
 **ignoreVerified** | **optional.Bool**| If true, return intercoms ignoring their verified value | 
 **recentlyChanged** | **optional.Bool**| If true, return all intercoms which recently were updated | 

### Return type

[**[]OpenerIntercomModel**](OpenerIntercomModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_2**
> OpenerIntercomModel Get_2(ctx, intercomId)
Get an intercom model



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **intercomId** | **int32**| The intercom ID | 

### Return type

[**OpenerIntercomModel**](OpenerIntercomModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

