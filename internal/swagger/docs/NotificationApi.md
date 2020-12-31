# \NotificationApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](NotificationApi.md#Delete) | **Delete** /notification/{notificationId} | Delete a notification configuration
[**Get**](NotificationApi.md#Get) | **Get** /notification | Get all notifications attached to your account
[**Get_0**](NotificationApi.md#Get_0) | **Get** /notification/{notificationId} | Get a notification configuration
[**Post**](NotificationApi.md#Post) | **Post** /notification/{notificationId} | Update a notification configuration
[**Put**](NotificationApi.md#Put) | **Put** /notification | Create a notification configuration


# **Delete**
> Delete(ctx, notificationId)
Delete a notification configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationId** | **string**| The unique notification ID | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> []Notification Get(ctx, optional)
Get all notifications attached to your account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceId** | **optional.String**| The reference ID to the third party system | 

### Return type

[**[]Notification**](Notification.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> Notification Get_0(ctx, notificationId)
Get a notification configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationId** | **string**| The unique notification ID | 

### Return type

[**Notification**](Notification.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Notification Post(ctx, body, notificationId)
Update a notification configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Notification**](Notification.md)| Notification update representation | 
  **notificationId** | **string**| The unique notification ID | 

### Return type

[**Notification**](Notification.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put**
> Notification Put(ctx, body)
Create a notification configuration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Notification**](Notification.md)| Notification representation | 

### Return type

[**Notification**](Notification.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

