# \AccountApi

All URIs are relative to *https://api.nuki.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](AccountApi.md#Delete) | **Delete** /account | Delete an account
[**Delete_0**](AccountApi.md#Delete_0) | **Delete** /account/otp | Disables one time password for an account
[**Delete_1**](AccountApi.md#Delete_1) | **Delete** /account/setting | Delete an account setting
[**Delete_2**](AccountApi.md#Delete_2) | **Delete** /account/sub/{accountId} | Delete an sub account
[**Get**](AccountApi.md#Get) | **Get** /account | Get an account
[**Get_0**](AccountApi.md#Get_0) | **Get** /account/setting | Get account setting
[**Get_1**](AccountApi.md#Get_1) | **Get** /account/sub | Get an list of sub accounts
[**Get_2**](AccountApi.md#Get_2) | **Get** /account/sub/{accountId} | Get an sub account
[**Post**](AccountApi.md#Post) | **Post** /account | Update an account
[**Post_0**](AccountApi.md#Post_0) | **Post** /account/otp | Enables one time password for an account
[**Post_1**](AccountApi.md#Post_1) | **Post** /account/password/reset | Reset account password
[**Post_2**](AccountApi.md#Post_2) | **Post** /account/sub/{accountId} | Update an sub account
[**Put**](AccountApi.md#Put) | **Put** /account | Create an account
[**Put_0**](AccountApi.md#Put_0) | **Put** /account/otp | Create an one time password secret
[**Put_1**](AccountApi.md#Put_1) | **Put** /account/setting | Create or update account setting
[**Put_2**](AccountApi.md#Put_2) | **Put** /account/sub | Create an sub account


# **Delete**
> Delete(ctx, )
Delete an account



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_0**
> Delete_0(ctx, )
Disables one time password for an account



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_1**
> Delete_1(ctx, )
Delete an account setting



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete_2**
> Delete_2(ctx, accountId)
Delete an sub account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**| The account id | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> MyAccount Get(ctx, )
Get an account



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MyAccount**](MyAccount.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> AccountSetting Get_0(ctx, )
Get account setting



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountSetting**](AccountSetting.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> Account Get_1(ctx, optional)
Get an list of sub accounts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGet_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGet_5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| The optional email (regex) | 

### Return type

[**Account**](Account.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_2**
> Account Get_2(ctx, accountId)
Get an sub account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int32**| The account id | 

### Return type

[**Account**](Account.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, body, optional)
Update an account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountUpdate**](AccountUpdate.md)| Account update representation | 
 **optional** | ***AccountApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteApiTokens** | **optional.Bool**| If false existing API tokens are not deleted if the password is changed | [default to true]

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, body)
Enables one time password for an account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountOtpEnable**](AccountOtpEnable.md)| Account one time password enable representation | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_1**
> Post_1(ctx, body)
Reset account password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountPasswordReset**](AccountPasswordReset.md)| Account password reset representation | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_2**
> Post_2(ctx, body, accountId)
Update an sub account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountSubUpdate**](AccountSubUpdate.md)| Account update representation | 
  **accountId** | **int32**| The account id | 

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
Create an account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountCreate**](AccountCreate.md)| Account create representation | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put_0**
> string Put_0(ctx, )
Create an one time password secret



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

# **Put_1**
> AccountSetting Put_1(ctx, body)
Create or update account setting



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountSetting**](AccountSetting.md)| Account setting representation | 

### Return type

[**AccountSetting**](AccountSetting.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Put_2**
> MyAccount Put_2(ctx, body)
Create an sub account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountSubCreate**](AccountSubCreate.md)| Account sub create representation | 

### Return type

[**MyAccount**](MyAccount.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

